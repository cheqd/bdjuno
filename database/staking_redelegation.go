package database

import (
	"fmt"

	dbtypes "github.com/forbole/bdjuno/v3/database/types"
	"github.com/forbole/bdjuno/v3/types"
)

// SaveRedelegations allows to store the given redelegations
func (db *Db) SaveRedelegations(redelegations []types.Redelegation) error {
	accQry := `
INSERT INTO account (address) VALUES `
	var accParams []interface{}

	rdQry := `
INSERT INTO redelegation 
    (delegator_address, src_validator_address, dst_validator_address, amount, completion_time, height) 
VALUES `
	var rdParams []interface{}

	for i, redelegation := range redelegations {
		a1 := i * 1
		accQry += fmt.Sprintf("($%d),", a1+1)
		accParams = append(accParams, redelegation.DelegatorAddress)

		// Get the validators info
		srcVal, err := db.GetValidator(redelegation.SrcValidator)
		if err != nil {
			return err
		}

		dstVal, err := db.GetValidator(redelegation.DstValidator)
		if err != nil {
			return err
		}

		// Convert the amount value
		coin := dbtypes.NewDbCoin(redelegation.Amount)
		amountValue, err := coin.Value()
		if err != nil {
			return err
		}

		rdi := i * 6
		rdQry += fmt.Sprintf("($%d,$%d,$%d,$%d,$%d,$%d),", rdi+1, rdi+2, rdi+3, rdi+4, rdi+5, rdi+6)
		rdParams = append(rdParams,
			redelegation.DelegatorAddress,
			srcVal.GetConsAddr(), dstVal.GetConsAddr(), amountValue, redelegation.CompletionTime, redelegation.Height)
	}

	// Insert the delegators
	accQry = accQry[:len(accQry)-1] // Remove the trailing ","
	accQry += " ON CONFLICT DO NOTHING"
	_, err := db.Sql.Exec(accQry, accParams...)
	if err != nil {
		return err
	}

	// Insert the redelegations
	rdQry = rdQry[:len(rdQry)-1] // Remove the trailing ","
	rdQry += `ON CONFLICT ON CONSTRAINT redelegation_validator_delegator_unique DO UPDATE 
	SET 
		amount = excluded.amount,
		height = excluded.height
	WHERE redelegation.height <= excluded.height`

	_, err = db.Sql.Exec(rdQry, rdParams...)
	return err
}

// DeleteRedelegation removes the given redelegation from the database
func (db *Db) DeleteRedelegation(redelegation types.Redelegation) error {
	srcVal, err := db.GetValidator(redelegation.SrcValidator)
	if err != nil {
		return err
	}

	dstVal, err := db.GetValidator(redelegation.DstValidator)
	if err != nil {
		return err
	}

	stmt := `
DELETE FROM redelegation 
WHERE delegator_address = $1 
  AND src_validator_address = $2 
  AND dst_validator_address = $3 
  AND completion_time = $4`
	_, err = db.Sql.Exec(stmt,
		redelegation.DelegatorAddress, srcVal.GetConsAddr(), dstVal.GetOperator(), redelegation.CompletionTime,
	)
	return err
}
