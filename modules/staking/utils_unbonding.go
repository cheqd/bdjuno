package staking

import (
	"time"

	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/forbole/bdjuno/v3/types"
	juno "github.com/forbole/juno/v3/types"
	"github.com/rs/zerolog/log"
)

// storeRedelegationFromMessage handles a MsgBeginRedelegate by saving the redelegation inside the database,
// and returns the new redelegation instance
func (m *Module) storeUnbondingDelegationFromMessage(
	tx *juno.Tx, index int, msg *stakingtypes.MsgUndelegate,
) (*types.UnbondingDelegation, error) {
	event, err := tx.FindEventByType(index, stakingtypes.EventTypeUnbond)
	if err != nil {
		return nil, err
	}

	completionTimeStr, err := tx.FindAttributeByKey(event, stakingtypes.AttributeKeyCompletionTime)
	if err != nil {
		return nil, err
	}

	completionTime, err := time.Parse(time.RFC3339, completionTimeStr)
	if err != nil {
		return nil, err
	}

	unbonding := types.NewUnbondingDelegation(
		msg.DelegatorAddress,
		msg.ValidatorAddress,
		msg.Amount,
		completionTime,
		tx.Height,
	)

	return &unbonding, m.db.SaveUnbondingDelegations([]types.UnbondingDelegation{unbonding})
}

// deleteUnbondingDelegation returns a function that when called deletes the given delegation from the database
func (m *Module) deleteUnbondingDelegation(delegation types.UnbondingDelegation) func() {
	return func() {
		err := m.db.DeleteUnbondingDelegation(delegation)
		if err != nil {
			log.Error().Str("module", "staking").Err(err).
				Str("operation", "delete unbonding delegation").Msg("error while deleting unbonding delegation")
		}
	}
}
