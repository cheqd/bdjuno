package utils

import "github.com/forbole/callisto/v4/types"

const (
	maxPostgreSQLParams = 65535
)

func SplitAccounts(accounts []types.Account, paramsNumber int) [][]types.Account {
	maxBalancesPerSlice := maxPostgreSQLParams / paramsNumber
	slices := make([][]types.Account, len(accounts)/maxBalancesPerSlice+1)

	sliceIndex := 0
	for index, account := range accounts {
		slices[sliceIndex] = append(slices[sliceIndex], account)

		if index > 0 && index%(maxBalancesPerSlice-1) == 0 {
			sliceIndex++
		}
	}

	return slices
}

func SplitTopAccounts(accounts []types.TopAccount, paramsNumber int) [][]types.TopAccount {
	maxBalancesPerSlice := maxPostgreSQLParams / paramsNumber
	slices := make([][]types.TopAccount, len(accounts)/maxBalancesPerSlice+1)

	sliceIndex := 0
	for index, account := range accounts {
		slices[sliceIndex] = append(slices[sliceIndex], account)

		if index > 0 && index%(maxBalancesPerSlice-1) == 0 {
			sliceIndex++
		}
	}

	return slices
}

func SplitBalances(balances []types.NativeTokenAmount, paramsNumber int) [][]types.NativeTokenAmount {
	maxPostgreSQLParams := 65535
	maxBalancesPerSlice := maxPostgreSQLParams / paramsNumber

	slices := make([][]types.NativeTokenAmount, len(balances)/maxBalancesPerSlice+1)

	sliceIndex := 0
	for index, balance := range balances {
		slices[sliceIndex] = append(slices[sliceIndex], balance)

		if index > 0 && index%(maxBalancesPerSlice-1) == 0 {
			sliceIndex++
		}
	}

	return slices
}
