package config

import (
	"fmt"
	add_account "github.com/doga10/clean-architecture-go/src/data/usecases/account/add-account"
	load_account_by_email "github.com/doga10/clean-architecture-go/src/data/usecases/account/load-account-by-email"
	account2 "github.com/doga10/clean-architecture-go/src/domain/usecases/account"
	bcrypt_adapter "github.com/doga10/clean-architecture-go/src/infra/criptography/bcrypt-adapter"
	"github.com/doga10/clean-architecture-go/src/infra/db/mongodb/account"
	"github.com/doga10/clean-architecture-go/src/infra/db/mongodb/helpers"
	"github.com/gin-gonic/gin"
)

func StartApp() {
	router := gin.Default()
	crypt := bcrypt_adapter.NewBcryptAdapter()
	collection := helpers.GetCollection("accounts")
	repo := account.NewAccountMongoRepository(collection)
	load_account_by_email.NewDbLoadAccountByEmail(repo)
	svc := add_account.NewDbAddAccount(repo, repo, crypt)
	var element account2.AddAccountParams
	element.Name = "Douglas Dennys"
	element.Email = "douglasdennys4@gmail.com"
	element.Password = "password"

	cur, err := svc.Add(&element)
	fmt.Println(err)
	fmt.Println(cur)

	err = router.Run()
	if err != nil {
		panic(err)
	}
}
