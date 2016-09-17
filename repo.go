package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var database Database

func init() {
	file, e := ioutil.ReadFile("./db.json")
	if e != nil {
		s := fmt.Sprintf("File error: %v\n", e)
		os.Stderr.Write([]byte(s))
		os.Exit(1)
	}

	json.Unmarshal(file, &database)

	pretty, err := json.MarshalIndent(database, "", " ")
	if err != nil {
		os.Stderr.Write([]byte("Cannot pretty print"))
		fmt.Printf("Database: %v\n", database)
	} else {
		os.Stdout.Write([]byte("Using the following database:\n"))
		os.Stdout.Write(pretty)
	}
}
