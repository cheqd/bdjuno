package utils

import (
	"fmt"

	"github.com/cheqd/juno/v4/node"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"
)

// QueryTxs queries all the transactions from the given node corresponding to the given query
func QueryTxs(node node.Node, query string) ([]*coretypes.ResultTx, error) {
	var txs []*coretypes.ResultTx

	page := 1
	perPage := 100
	stop := false
	for !stop {
		result, err := node.TxSearch(query, &page, &perPage, "")
		if err != nil {
			return nil, fmt.Errorf("error while running tx search: %s", err)
		}

		page++
		txs = append(txs, result.Txs...)
		stop = len(txs) == result.TotalCount
	}

	return txs, nil
}
