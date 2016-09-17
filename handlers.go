package main

import (
	"encoding/json"
	"net/http"
	"sort"
	"strconv"
)

func BalanceList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	result := make(Balances, len(database.BalancesDB))
	copy(result, database.BalancesDB)
	sort.Sort(sort.Reverse(BalancesByDate(result)))

	if err := json.NewEncoder(w).Encode(result); err != nil {
		panic(err)
	}
}

func BalanceQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	balances := database.BalancesDB

	account := r.URL.Query().Get("account")
	dateFrom, dateFromExists := parseDate(r.URL.Query().Get("dateFrom"))
	dateTo, dateToExists := parseDate(r.URL.Query().Get("dateTo"))

	result := make(Balances, 0)
	for _, balance := range balances {
		accountOK := account == "" || balance.Account == account
		dateFromOK := !dateFromExists || dateFrom <= balance.Timestamp
		dateToOK := !dateToExists || dateTo >= balance.Timestamp

		if accountOK && dateFromOK && dateToOK {
			result = append(result, balance)
		}
	}

	sort.Sort(sort.Reverse(BalancesByDate(result)))

	if err := json.NewEncoder(w).Encode(result); err != nil {
		panic(err)
	}
}

func AccountList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	result := make(Accounts, len(database.AccountsDB))
	copy(result, database.AccountsDB)
	sort.Sort(AccountsByName(result))

	if err := json.NewEncoder(w).Encode(result); err != nil {
		panic(err)
	}
}

func AccountQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	accounts := database.AccountsDB

	acct := r.URL.Query().Get("account")
	name := r.URL.Query().Get("name")
	currency := r.URL.Query().Get("currency")

	result := make([]Account, 0)
	for _, account := range accounts {
		accountOK := acct == "" || account.Account == acct
		nameOK := name == "" || account.Name == name
		currencyOK := currency == "" || account.Currency == currency

		if accountOK && nameOK && currencyOK {
			result = append(result, account)
		}
	}

	sort.Sort(AccountsByName(result))

	if err := json.NewEncoder(w).Encode(result); err != nil {
		panic(err)
	}
}

func parseDate(strMillis string) (int64, bool) {
	if strMillis == "" {
		return 0, false
	}

	if millis, err := strconv.ParseInt(strMillis, 10, 64); err != nil {
		return 0, false
	} else {
		return millis, true
	}
}
