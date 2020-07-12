package add_account

import (
	usecases "github.com/doga10/clean-architecture-go/src/main/factories/usecases/account/add-account"
	add "github.com/doga10/clean-architecture-go/src/presentation/controllers/account/add-account"
	"github.com/doga10/clean-architecture-go/src/presentation/protocols"
)

func MakeAddAccountController() protocols.Controller {
	return add.NewAddAccountController(usecases.MakeDbAddAccount(), MakeAddAccountValidation())
}