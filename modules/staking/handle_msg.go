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

	case *stakingtypes.MsgDelegate:
		return m.storeDelegationFromMessage(tx.Height, cosmosMsg)

	case *stakingtypes.MsgBeginRedelegate:
		return m.handleMsgBeginRedelegate(tx, index, cosmosMsg)

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

// handleMsgBeginRedelegate handles a MsgBeginRedelegate storing the data inside the database
func (m *Module) handleMsgBeginRedelegate(
	tx *juno.Tx, index int, msg *stakingtypes.MsgBeginRedelegate) error {
	redelegation, err := m.storeRedelegationFromMessage(tx, index, msg)
	if err != nil {
		return err
	}

	// When the time expires, update the delegations and delete this redelegation
	time.AfterFunc(time.Until(redelegation.CompletionTime),
		m.refreshDelegations(tx.Height, msg.DelegatorAddress))
	time.AfterFunc(time.Until(redelegation.CompletionTime),
		m.deleteRedelegation(*redelegation))

	// Update the current delegations
	return m.updateDelegationsAndReplaceExisting(tx.Height, msg.DelegatorAddress)
}

// handleMsgUndelegate handles a MsgUndelegate storing the data inside the database
func (m *Module) handleMsgUndelegate(tx *juno.Tx, index int, msg *stakingtypes.MsgUndelegate) error {
	delegation, err := m.storeUnbondingDelegationFromMessage(tx, index, msg)
	if err != nil {
		return err
	}

	// When timer expires update the delegations, update the user balance and remove the unbonding delegation
	time.AfterFunc(time.Until(delegation.CompletionTimestamp),
		m.refreshDelegations(tx.Height, msg.DelegatorAddress))
	time.AfterFunc(time.Until(delegation.CompletionTimestamp),
		m.updateBalances(msg.DelegatorAddress))
	time.AfterFunc(time.Until(delegation.CompletionTimestamp),
		m.deleteUnbondingDelegation(*delegation))

	// Update the current delegations
	return m.updateDelegationsAndReplaceExisting(tx.Height, msg.DelegatorAddress)
}

func (m *Module) updateBalances(address string) func() {
	return func() {
		height, err := m.db.GetLastBlockHeight()
		if err != nil {
			log.Error().Err(err).Str("module", "bank").
				Str("operation", "refresh balance").Msg("error while getting latest block height")
			return
		}

		err = m.bankModule.UpdateBalances([]string{address}, height)
		if err != nil {
			log.Error().Err(err).Str("module", "bank").
				Str("operation", "refresh balance").Msg("error while updating balance")
		}
	}
}
