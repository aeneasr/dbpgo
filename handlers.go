package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func BalanceList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(database.BalancesDB); err != nil {
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

	result := make([]Balance, 0)
	for _, balance := range balances {
		accountOK := account == "" || balance.Account == account
		dateFromOK := !dateFromExists || dateFrom <= balance.Timestamp
		dateToOK := !dateToExists || dateTo >= balance.Timestamp

		if accountOK && dateFromOK && dateToOK {
			result = append(result, balance)
		}
	}

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
