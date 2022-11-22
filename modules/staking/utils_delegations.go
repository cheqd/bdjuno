package staking

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/types/query"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/forbole/bdjuno/v3/types"
	"github.com/rs/zerolog/log"
)

// storeDelegationFromMessage handles a MsgDelegate and saves the delegation inside the database
func (m *Module) storeDelegationFromMessage(
	height int64, msg *stakingtypes.MsgDelegate,
) error {
	res, err := m.source.GetDelegation(height, msg.DelegatorAddress, msg.ValidatorAddress)
	if err != nil {
		return fmt.Errorf("error while getting delegator delegations: %s", err)
	}

	delegation := types.NewDelegation(
		res.DelegationResponse.Delegation.DelegatorAddress,
		res.DelegationResponse.Delegation.ValidatorAddress,
		res.DelegationResponse.Balance,
		height,
	)

	return m.db.SaveDelegations([]types.Delegation{delegation})
}

// refreshDelegations returns a function that when called updates the delegations of the provided delegator.
// In order to properly update the data, it removes all the existing delegations and stores new ones querying the gRPC
func (m *Module) refreshDelegations(height int64, delegator string) func() {
	return func() {
		err := m.updateDelegationsAndReplaceExisting(height, delegator)
		if err != nil {
			log.Error().Str("module", "staking").Err(err).
				Str("operation", "refresh delegations").Msg("error while refreshing delegations")
		}
	}
}

// updateDelegationsAndReplaceExisting updates the delegations of the given delegator by querying them at the
// required height, and then stores them inside the database by replacing all existing ones.
func (m *Module) updateDelegationsAndReplaceExisting(height int64, delegator string) error {
	// Remove existing delegations
	err := m.db.DeleteDelegatorDelegations(delegator)
	if err != nil {
		return err
	}

	return m.updateDelegations(height, delegator)
}

// updateDelegations updates the current delegations for the given delegator by removing all the existing ones and
// getting the new ones from the client
func (m *Module) updateDelegations(height int64, delegator string) error {
	// Get the delegations
	var delRes []stakingtypes.DelegationResponse
	var nextKey []byte
	stop := false
	for !stop {
		res, err := m.source.GetDelegationsWithPagination(
			height, delegator, &query.PageRequest{},
		)
		if err != nil {
			return err
		}
		nextKey = res.Pagination.NextKey
		stop = len(nextKey) == 0
		delRes = append(delRes, res.DelegationResponses...)
	}

	var delegations = make([]types.Delegation, len(delRes))
	for index, r := range delRes {
		delegations[index] = types.NewDelegation(
			r.Delegation.DelegatorAddress,
			r.Delegation.ValidatorAddress,
			r.Balance,
			height,
		)
	}

	return m.db.SaveDelegations(delegations)
}

// deleteRedelegation returns a function that when called removes the given redelegation from the database.
func (m *Module) deleteRedelegation(redelegation types.Redelegation) func() {
	return func() {
		// Remove existing redelegations
		err := m.db.DeleteRedelegation(redelegation)
		if err != nil {
			log.Error().Str("module", "staking").Err(err).
				Str("operation", "update redelegations").
				Msg("error while removing delegator redelegations")
			return
		}
	}
}
