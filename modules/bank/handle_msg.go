package bank

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
	"github.com/rs/zerolog/log"

	"github.com/forbole/bdjuno/v3/modules/utils"

	juno "github.com/forbole/juno/v3/types"
)

// HandleMsg handles any message updating the involved addresses balances
func (m *Module) HandleMsg(index int, msg sdk.Msg, tx *juno.Tx) error {
	addresses, err := m.messageParser(m.cdc, msg)
	if err != nil {
		log.Error().Str("module", "bank").Err(err).
			Str("operation", "refresh balances").
			Msgf("error while refreshing balances after message of type %s", proto.MessageName(msg))
	}

	return m.UpdateBalances(utils.FilterNonAccountAddresses(addresses), tx.Height)
}
