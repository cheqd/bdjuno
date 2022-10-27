package types

// FeeAllowanceRow represents a single row inside the fee_grant_allowance table
type FeeAllowanceRow struct {
	Grantee   string `db:"grantee_address"`
	Granter   string `db:"granter_address"`
	Allowance string `db:"allowance"`
	ID        uint64 `db:"id"`
	Height    int64  `db:"height"`
}
