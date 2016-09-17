package main

type Database struct {
	BalancesDB Balances `json:"balances"`
	AccountsDB Accounts `json:"accounts"`
}
