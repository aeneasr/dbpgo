package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	Queries     []string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Queries(route.Queries...).
			Handler(route.HandlerFunc)

	}

	return router
}

var routes = Routes{
	Route{
		"BalanceQuery",
		"GET",
		"/balances/query",
		[]string{},
		BalanceQuery,
	},
	Route{
		"BalanceList",
		"GET",
		"/balances",
		[]string{},
		BalanceList,
	},
}
