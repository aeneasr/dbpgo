package main

type Balance struct {
	Id          int     `json:"id"`
	Account     string  `json:"account"`
	AccountName string  `json:"accountName"`
	Currency    string  `json:"currency"`
	Amount      float64 `json:"amount"`
	Timestamp   int64   `json:"date"`
}

type Balances []Balance

type Database struct {
	BalancesDB []Balance `json:"balances"`
}
