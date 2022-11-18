package bank

import (
	"encoding/json"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/forbole/bdjuno/v3/types"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/rs/zerolog/log"
)

// HandleGenesis implements modules.GenesisModule
func (m *Module) HandleGenesis(doc *tmtypes.GenesisDoc, appState map[string]json.RawMessage) error {
	log.Debug().Str("module", "bank").Msg("parsing genesis")

	var bankState banktypes.GenesisState
	if err := m.cdc.UnmarshalJSON(appState[banktypes.ModuleName], &bankState); err != nil {
		return err
	}

	// Store the accounts
	balances := make([]types.AccountBalance, len(bankState.Balances))
	for index, balance := range bankState.Balances {
		balances[index] = types.NewAccountBalance(balance.Address, balance.Coins, doc.InitialHeight)
	}

	return m.db.SaveAccountBalances(balances)
}
