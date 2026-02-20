package types

import (
	"fmt"
	"reflect"

	"github.com/forbole/juno/v6/parser"

	nodebuilder "github.com/forbole/juno/v6/node/builder"
	"github.com/forbole/juno/v6/types/config"

	"github.com/forbole/juno/v6/database"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"

	modsregistrar "github.com/forbole/juno/v6/modules/registrar"
)

// GetParserContext setups all the things that can be used to later parse the chain state
func GetParserContext(cfg config.Config, parseConfig *Config) (*parser.Context, error) {
	// Setup the SDK configuration
	sdkConfig, sealed := getConfig()
	if !sealed {
		parseConfig.GetSetupConfig()(cfg, sdkConfig)
		sdkConfig.Seal()
	}

	// Get the db
	databaseCtx := database.NewContext(cfg.Database, parseConfig.GetLogger())
	db, err := parseConfig.GetDBBuilder()(databaseCtx)
	if err != nil {
		return nil, err
	}

	// Init the client
	// Create txConfig for remote node to decode transactions
	cdc := parseConfig.GetCodec()
	var txConfig client.TxConfig
	if cdc != nil {
		txConfig = authtx.NewTxConfig(cdc, authtx.DefaultSignModes)
	}
	cp, err := nodebuilder.BuildNode(cfg.Node, txConfig, cdc)
	if err != nil {
		return nil, fmt.Errorf("failed to start client: %s", err)
	}

	// Setup the logging
	err = parseConfig.GetLogger().SetLogFormat(cfg.Logging.LogFormat)
	if err != nil {
		return nil, fmt.Errorf("error while setting logging format: %s", err)
	}

	err = parseConfig.GetLogger().SetLogLevel(cfg.Logging.LogLevel)
	if err != nil {
		return nil, fmt.Errorf("error while setting logging level: %s", err)
	}

	// Get the modules
	context := modsregistrar.NewContext(cfg, sdkConfig, db, cp, parseConfig.GetLogger())
	mods := parseConfig.GetRegistrar().BuildModules(context)
	registeredModules := modsregistrar.GetModules(mods, cfg.Chain.Modules, parseConfig.GetLogger())

	return parser.NewContext(cp, db, parseConfig.GetLogger(), registeredModules), nil
}

// getConfig returns the SDK Config instance as well as if it's sealed or not
func getConfig() (config *sdk.Config, sealed bool) {
	sdkConfig := sdk.GetConfig()
	fv := reflect.ValueOf(sdkConfig).Elem().FieldByName("sealed")
	return sdkConfig, fv.Bool()
}
