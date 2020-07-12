package main

import (
	"github.com/doga10/clean-architecture-go/src/infra/db/mongodb/helpers"
	"github.com/doga10/clean-architecture-go/src/main/config"
)

func main() {
	err := helpers.Connect("mongodb://127.0.0.1:27017", "app")
	if err != nil {
		panic(err)
	}

	config.StartApp()
}
