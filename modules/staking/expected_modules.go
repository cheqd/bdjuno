package staking

type BankModule interface {
	UpdateBalances(addresses []string, height int64) error
}
