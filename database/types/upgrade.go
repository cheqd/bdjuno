package types

type SoftwareUpgradePlanRow struct {
	PlanName      string `db:"plan_name"`
	Info          string `db:"info"`
	ProposalID    uint64 `db:"proposal_id"`
	UpgradeHeight int64  `db:"upgrade_height"`
	Height        int64  `db:"height"`
}

func NewSoftwareUpgradePlanRow(
	proposalID uint64, planName string, upgradeHeight int64, info string, height int64,
) SoftwareUpgradePlanRow {
	return SoftwareUpgradePlanRow{
		ProposalID:    proposalID,
		PlanName:      planName,
		UpgradeHeight: upgradeHeight,
		Info:          info,
		Height:        height,
	}
}
