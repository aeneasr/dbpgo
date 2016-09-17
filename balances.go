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

type BalancesByDate Balances

func (b BalancesByDate) Len() int {
	return len(b)
}

func (b BalancesByDate) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b BalancesByDate) Less(i, j int) bool {
	return b[i].Timestamp < b[j].Timestamp
}
