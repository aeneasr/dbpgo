package main

type Account struct {
	Id       string `json:"id"`
	Account  string `json:"account"`
	Name     string `json:"name"`
	Currency string `json:"currency"`
	Country  string `json:"country"`
	Address  string `json:"address"`
}

type Accounts []Account

type AccountsByName Accounts

func (a AccountsByName) Len() int {
	return len(a)
}

func (a AccountsByName) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a AccountsByName) Less(i, j int) bool {
	return a[i].Name < a[j].Name
}
