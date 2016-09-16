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
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	json.Unmarshal(file, &database)
	fmt.Printf("Database: %v\n", database)
}
