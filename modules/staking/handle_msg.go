package staking

import (
	"fmt"
	"time"

	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/rs/zerolog/log"

	sdk "github.com/cosmos/cosmos-sdk/types"
	juno "github.com/forbole/juno/v3/types"
)

// HandleMsg implements MessageModule
func (m *Module) HandleMsg(index int, msg sdk.Msg, tx *juno.Tx) error {
	if len(tx.Logs) == 0 {
		return nil
	}

	switch cosmosMsg := msg.(type) {
	case *stakingtypes.MsgCreateValidator:
		return m.handleMsgCreateValidator(tx.Height, cosmosMsg)

	case *stakingtypes.MsgEditValidator:
		return m.handleEditValidator(tx.Height, cosmosMsg)

	case *stakingtypes.MsgUndelegate:
		return m.handleMsgUndelegate(tx, index, cosmosMsg)
	}

	return nil
}

// ---------------------------------------------------------------------------------------------------------------------

// handleMsgCreateValidator handles properly a MsgCreateValidator instance by
// saving into the database all the data associated to such validator
func (m *Module) handleMsgCreateValidator(height int64, msg *stakingtypes.MsgCreateValidator) error {
	err := m.RefreshValidatorInfos(height, msg.ValidatorAddress)
	if err != nil {
		return fmt.Errorf("error while refreshing validator from MsgCreateValidator: %s", err)
	}
	return nil
}

// handleEditValidator handles MsgEditValidator utils, updating the validator info
func (m *Module) handleEditValidator(height int64, msg *stakingtypes.MsgEditValidator) error {
	err := m.RefreshValidatorInfos(height, msg.ValidatorAddress)
	if err != nil {
		return fmt.Errorf("error while refreshing validator from MsgEditValidator: %s", err)
	}

	return nil
}

// handleMsgUndelegate handles a MsgUndelegate storing the data inside the database
func (m *Module) handleMsgUndelegate(
	tx *juno.Tx, index int, msg *stakingtypes.MsgUndelegate) error {
	event, err := tx.FindEventByType(index, stakingtypes.EventTypeUnbond)
	if err != nil {
		return err
	}

	completionTimeStr, err := tx.FindAttributeByKey(event, stakingtypes.AttributeKeyCompletionTime)
	if err != nil {
		return err
	}

	completionTime, err := time.Parse(time.RFC3339, completionTimeStr)
	if err != nil {
		return err
	}

	// When timer expires update the delegations, update the user balance
	time.AfterFunc(time.Until(completionTime),
		m.refreshBalance(msg.DelegatorAddress, tx.Height))

	// Update the current delegations
	return nil
}

// refreshBalance returns a function that when called refreshes the balance of the user having the given address
func (m *Module) refreshBalance(address string, height int64) func() {
	return func() {
		err := m.bankModule.UpdateBalances([]string{address}, height)
		if err != nil {
			log.Error().Err(err).Str("module", "bank").
				Str("operation", "refresh balance").Msg("error while updating balance")
		}
	}
}
