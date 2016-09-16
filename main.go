package main

import (
	"github.com/gorilla/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS()(router)))
}
