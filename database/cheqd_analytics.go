package database

import (
	"fmt"

	dbtypes "github.com/forbole/bdjuno/v3/database/types"
)

func (db *Db) GetTotalDIDs(msgTypes []string, plottable bool) ([]dbtypes.AnalyticsItem, error) {
	var buildTypes string
	for i, msg := range msgTypes {
		if i == 0 {
			buildTypes += fmt.Sprintf("type='cheqdid.cheqdnode.cheqd.v1.%s'", msg)
			continue
		}
		buildTypes += fmt.Sprintf("or type='cheqdid.cheqdnode.cheqd.v1.%s'", msg)
	}

	// keep the length at least one so that we can cover non-plottable case where we only fill a single item
	dids := make([]dbtypes.AnalyticsItem, 1)
	if plottable {
		query := fmt.Sprintf("SELECT COUNT(*) AS total, height from message WHERE %s GROUP BY height", buildTypes)
		err := db.Sqlx.Select(&dids, query)
		if err != nil {
			return nil, err
		}

		return dids, nil
	}

	query := `SELECT COUNT(*) AS total from message WHERE ` + buildTypes
	err := db.Sqlx.QueryRow(query).Scan(&dids[0].Total)
	if err != nil {
		return nil, err
	}

	return dids, nil
}

func (db *Db) GetTotalResources(msgTypes []string, plottable bool) ([]dbtypes.AnalyticsItem, error) {
	var buildTypes string
	for i, msg := range msgTypes {
		if i == 0 {
			buildTypes += fmt.Sprintf("like cheqdid.cheqdnode.resource.%s", msg)
			continue
		}
		buildTypes += fmt.Sprintf("or like cheqdid.cheqdnode.resource.%s", msg)
	}

	// keep the length at least one so that we can cover non-plottable case where we only fill a single item
	resources := make([]dbtypes.AnalyticsItem, 1)
	if plottable {
		query := `SELECT COUNT(*) AS total, height from message WHERE type ` + buildTypes
		query += " GROUP BY height;"
		err := db.Sqlx.Select(&resources, query)
		if err != nil {
			return nil, err
		}

		return resources, nil
	}

	query := `SELECT COUNT(*) AS total from message WHERE ` + buildTypes
	err := db.Sqlx.QueryRow(query).Scan(&resources[0].Total)
	if err != nil {
		return nil, err
	}

	return resources, nil
}
