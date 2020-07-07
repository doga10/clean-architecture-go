package config

import (
	"fmt"
	add_account "github.com/doga10/clean-architecture-go/src/data/usecases/account/add-account"
	account2 "github.com/doga10/clean-architecture-go/src/domain/usecases/account"
	"github.com/doga10/clean-architecture-go/src/infra/db/mongodb/account"
	"github.com/doga10/clean-architecture-go/src/infra/db/mongodb/helpers"
)

func StartApp() {
	collection := helpers.GetCollection("accounts")
	repo := account.NewAccountMongoRepository(collection)
	svc := add_account.NewDbAddAccount(repo)
	var element account2.AddAccountParams

	element.Name = "Douglas Dennys"
	element.Email = "douglasdennys45@gmail.com"
	element.Password = "123"

	cur, err := svc.Add(&element)
	fmt.Println(err)
	fmt.Println(cur)
}