package main

import (
	"github.com/doga10/clean-architecture-go/src/infra/db/mongodb/helpers"
	"github.com/doga10/clean-architecture-go/src/main/config"
	"os"
)

func main() {
	err := helpers.Connect(os.Getenv("MONGO_URL"), os.Getenv("MONGO_NAME"))
	if err != nil {
		panic(err)
	}

	config.StartApp()
}
