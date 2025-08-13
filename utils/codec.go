package utils

import (
	"sync"

	"cosmossdk.io/x/circuit"
	"cosmossdk.io/x/evidence"
	feegrantmodule "cosmossdk.io/x/feegrant/module"
	"cosmossdk.io/x/tx/signing"
	"cosmossdk.io/x/upgrade"
	app "github.com/cheqd/cheqd-node/app"
	did "github.com/cheqd/cheqd-node/x/did"
	"github.com/cheqd/cheqd-node/x/resource"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/address"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/std"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	authzmodule "github.com/cosmos/cosmos-sdk/x/authz/module"
	"github.com/cosmos/cosmos-sdk/x/bank"
	consensus "github.com/cosmos/cosmos-sdk/x/consensus"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	groupmodule "github.com/cosmos/cosmos-sdk/x/group/module"
	"github.com/cosmos/cosmos-sdk/x/mint"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/gogoproto/proto"
	icq "github.com/cosmos/ibc-apps/modules/async-icq/v8"
	"github.com/cosmos/ibc-go/modules/capability"
	ica "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts"
	ibcfee "github.com/cosmos/ibc-go/v8/modules/apps/29-fee"
	transfer "github.com/cosmos/ibc-go/v8/modules/apps/transfer"
	ibc "github.com/cosmos/ibc-go/v8/modules/core"
	ibctm "github.com/cosmos/ibc-go/v8/modules/light-clients/07-tendermint"
	feeabsmodule "github.com/osmosis-labs/fee-abstraction/v8/x/feeabs"
	feemarketmodule "github.com/skip-mev/feemarket/x/feemarket"
)

var (
	once sync.Once
	cdc  *codec.ProtoCodec
)

func GetCodec() codec.Codec {
	once.Do(func() {
		signopts := signing.Options{
			AddressCodec: address.Bech32Codec{
				Bech32Prefix: app.AccountAddressPrefix,
			},
			ValidatorAddressCodec: address.Bech32Codec{
				Bech32Prefix: app.ValidatorAddressPrefix,
			},
		}
		interfaceRegistry, _ := codectypes.NewInterfaceRegistryWithOptions(
			codectypes.InterfaceRegistryOptions{
				ProtoFiles:     proto.HybridResolver,
				SigningOptions: signopts,
			},
		)
		getBasicManagers().RegisterInterfaces(interfaceRegistry)
		std.RegisterInterfaces(interfaceRegistry)
		cdc = codec.NewProtoCodec(interfaceRegistry)
	})
	return cdc
}

// getBasicManagers returns the various basic managers that are used to register the encoding to
// support custom messages.
// This should be edited by custom implementations if needed.
func getBasicManagers() module.BasicManager {
	return module.NewBasicManager(
		auth.AppModuleBasic{},
		genutil.NewAppModuleBasic(genutiltypes.DefaultMessageValidator),
		bank.AppModuleBasic{},
		capability.AppModuleBasic{},
		staking.AppModuleBasic{},
		mint.AppModuleBasic{},
		distr.AppModuleBasic{},
		gov.NewAppModuleBasic([]govclient.ProposalHandler{
			paramsclient.ProposalHandler,
			feeabsmodule.UpdateAddHostZoneClientProposalHandler,
			feeabsmodule.UpdateDeleteHostZoneClientProposalHandler,
			feeabsmodule.UpdateSetHostZoneClientProposalHandler,
		}),
		params.AppModuleBasic{},
		crisis.AppModuleBasic{},
		slashing.AppModuleBasic{},
		ibc.AppModuleBasic{},
		ibctm.AppModuleBasic{},
		transfer.AppModuleBasic{},
		ica.AppModuleBasic{},
		feegrantmodule.AppModuleBasic{},
		upgrade.AppModuleBasic{},
		icq.AppModuleBasic{},
		evidence.AppModuleBasic{},
		authzmodule.AppModuleBasic{},
		groupmodule.AppModuleBasic{},
		vesting.AppModuleBasic{},
		consensus.AppModuleBasic{},
		circuit.AppModuleBasic{},
		did.AppModuleBasic{},
		resource.AppModuleBasic{},
		ibcfee.AppModuleBasic{},
		feeabsmodule.AppModuleBasic{},
		feemarketmodule.AppModuleBasic{},
	)
}

// UnpackMessage unpacks a message from a byte slice
func UnpackMessage[T proto.Message](cdc codec.Codec, bz []byte, ptr T) T {
	var any codectypes.Any
	cdc.MustUnmarshalJSON(bz, &any)
	var cosmosMsg sdk.Msg
	if err := cdc.UnpackAny(&any, &cosmosMsg); err != nil {
		panic(err)
	}
	return cosmosMsg.(T)
}
