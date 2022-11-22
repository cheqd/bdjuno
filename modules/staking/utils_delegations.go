package staking

import (
	"fmt"

	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/forbole/bdjuno/v3/types"
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
