package database

import (
	"fmt"

	dbtypes "github.com/forbole/bdjuno/v3/database/types"
	"github.com/forbole/bdjuno/v3/types"
)

// SaveUnbondingDelegations allows to store the given unbonding delegations
func (db *Db) SaveUnbondingDelegations(delegations []types.UnbondingDelegation) error {
	accQry := `
INSERT INTO account (address) VALUES `
	var accParams []interface{}

	udQry := `
INSERT INTO unbonding_delegation (validator_address, delegator_address, amount, completion_timestamp, height)
VALUES `
	var udParams []interface{}

	for i, delegation := range delegations {
		ai := i * 1
		accQry += fmt.Sprintf("($%d),", ai+1)
		accParams = append(accParams, delegation.DelegatorAddress)

		validator, err := db.GetValidator(delegation.ValidatorOperAddr)
		if err != nil {
			return err
		}

		coin := dbtypes.NewDbCoin(delegation.Amount)
		amount, err := coin.Value()
		if err != nil {
			return err
		}

		udi := i * 5
		udQry += fmt.Sprintf("($%d,$%d,$%d,$%d,$%d),", udi+1, udi+2, udi+3, udi+4, udi+5)
		udParams = append(udParams,
			validator.GetConsAddr(), delegation.DelegatorAddress, amount, delegation.CompletionTimestamp, delegation.Height)
	}

	// Insert the delegators
	accQry = accQry[:len(accQry)-1] // Remove the trailing ","
	accQry += " ON CONFLICT DO NOTHING"
	_, err := db.Sql.Exec(accQry, accParams...)
	if err != nil {
		return err
	}

	// Insert the current unbonding delegations
	udQry = udQry[:len(udQry)-1] // Remove the trailing ","
	udQry += ` ON CONFLICT ON CONSTRAINT unbonding_delegation_validator_delegator_unique DO UPDATE 
	SET 
		amount = excluded.amount,
		height = excluded.height
	WHERE unbonding_delegation.height <= excluded.height`

	_, err = db.Sql.Exec(udQry, udParams...)
	return err
}

// DeleteUnbondingDelegation removes the given unbonding delegation from the database
func (db *Db) DeleteUnbondingDelegation(delegation types.UnbondingDelegation) error {
	val, err := db.GetValidator(delegation.ValidatorOperAddr)
	if err != nil {
		return err
	}

	stmt := `
DELETE FROM unbonding_delegation 
WHERE delegator_address = $1 
  AND validator_address = $2 
  AND completion_timestamp = $3`
	_, err = db.Sql.Exec(stmt,
		delegation.DelegatorAddress, val.GetConsAddr(), delegation.CompletionTimestamp,
	)
	return err
}
