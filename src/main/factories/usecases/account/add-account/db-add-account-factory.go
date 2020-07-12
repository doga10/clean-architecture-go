package add_account

import (
	add_account "github.com/doga10/clean-architecture-go/src/data/usecases/account/add-account"
	"github.com/doga10/clean-architecture-go/src/domain/usecases/account"
	bcrypt_adapter "github.com/doga10/clean-architecture-go/src/infra/criptography/bcrypt-adapter"
	account2 "github.com/doga10/clean-architecture-go/src/infra/db/mongodb/account"
	"github.com/doga10/clean-architecture-go/src/infra/db/mongodb/helpers"
)

func MakeDbAddAccount() account.AddAccount {
	crypt := bcrypt_adapter.NewBcryptAdapter()
	collection := helpers.GetCollection("accounts")
	repo := account2.NewAccountMongoRepository(collection)

	return add_account.NewDbAddAccount(repo, repo, crypt)
}
