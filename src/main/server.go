package main

import (
	"fmt"
	"github.com/doga10/clean-architecture-go/src/infra/db/mongodb/helpers"
	"github.com/doga10/clean-architecture-go/src/main/config"
	"net/http"
)

func main() {
	var db *helpers.MongoDB
	err := db.Connect("mongodb://localhost:27017", "app")
	if err != nil {
		panic(err)
	}

	config.StartApp()

	fmt.Println("Successfully starting Go Lang service on host http://localhost:9090")
	err = http.ListenAndServe(":9090", nil)
	if err != nil {
		panic(err)
	}
}
