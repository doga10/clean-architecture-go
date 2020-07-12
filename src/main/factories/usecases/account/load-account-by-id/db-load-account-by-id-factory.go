package load_account_by_id

import (
	"github.com/doga10/clean-architecture-go/src/data/usecases/account/load-account-by-id"
	"github.com/doga10/clean-architecture-go/src/domain/usecases/account"
	db "github.com/doga10/clean-architecture-go/src/infra/db/mongodb/account"
	"github.com/doga10/clean-architecture-go/src/infra/db/mongodb/helpers"
)

func MakeLoadAccountById() account.LoadAccountById {
	collection := helpers.GetCollection("accounts")
	repo := db.NewAccountMongoRepository(collection)
	return load_account_by_id.NewDbLoadAccountById(repo)
}
