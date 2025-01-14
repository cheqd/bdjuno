package gov

import (
	tmctypes "github.com/cometbft/cometbft/rpc/core/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/forbole/callisto/v4/types"
)

type DistrModule interface {
	UpdateParams(height int64) error
}

type MintModule interface {
	UpdateParams(height int64) error
	UpdateInflation() error
}

type SlashingModule interface {
	UpdateParams(height int64) error
}

type StakingModule interface {
	GetStakingPool(height int64) (*types.Pool, error)
	GetStakingPoolSnapshot(height int64) (*types.PoolSnapshot, error)
	GetValidatorsWithStatus(height int64, status string) ([]stakingtypes.Validator, []types.Validator, error)
	GetValidatorsVotingPowers(height int64, vals *tmctypes.ResultValidators) ([]types.ValidatorVotingPower, error)
	GetValidatorsStatuses(height int64, validators []stakingtypes.Validator) ([]types.ValidatorStatus, error)
	UpdateParams(height int64) error
}
