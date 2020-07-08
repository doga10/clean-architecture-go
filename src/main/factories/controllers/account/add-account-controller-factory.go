package account

import (
	add_account "github.com/doga10/clean-architecture-go/src/data/usecases/account/add-account"
	load_account_by_email "github.com/doga10/clean-architecture-go/src/data/usecases/account/load-account-by-email"
	bcrypt_adapter "github.com/doga10/clean-architecture-go/src/infra/criptography/bcrypt-adapter"
	account2 "github.com/doga10/clean-architecture-go/src/infra/db/mongodb/account"
	"github.com/doga10/clean-architecture-go/src/infra/db/mongodb/helpers"
	"github.com/doga10/clean-architecture-go/src/presentation/controllers/account"
	"github.com/doga10/clean-architecture-go/src/presentation/protocols"
)

func MakeAddAccountController() protocols.Controller {
	crypt := bcrypt_adapter.NewBcryptAdapter()
	collection := helpers.GetCollection("accounts")
	repo := account2.NewAccountMongoRepository(collection)
	load_account_by_email.NewDbLoadAccountByEmail(repo)
	svc := add_account.NewDbAddAccount(repo, repo, crypt)

	return account.NewAddAccountController(svc)
}