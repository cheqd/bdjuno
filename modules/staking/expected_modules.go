package staking

type BankdModule interface {
	UpdateBalances(addresses []string, height int64) error
}
