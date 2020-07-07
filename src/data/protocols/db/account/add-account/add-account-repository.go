package add_account

import (
	"github.com/doga10/clean-architecture-go/src/domain/usecases/account"
)

type AddAccountRepository interface {
	Add(accountData *account.AddAccountParams) (interface{}, error)
}
