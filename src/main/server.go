package main

import (
	"github.com/doga10/clean-architecture-go/src/infra/db/mongodb/helpers"
	"github.com/doga10/clean-architecture-go/src/main/config"
)

func main() {
	var db *helpers.MongoDB
	err := db.Connect("mongodb://localhost:27017", "app")
	if err != nil {
		panic(err)
	}

	config.StartApp()
}
