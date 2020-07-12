package load_account_by_id

import (
	usecases "github.com/doga10/clean-architecture-go/src/main/factories/usecases/account/load-account-by-id"
	controller "github.com/doga10/clean-architecture-go/src/presentation/controllers/account/load-account-by-id"
	"github.com/doga10/clean-architecture-go/src/presentation/protocols"
)

func MakeLoadAccountByIdController() protocols.Controller {
	return controller.NewLoadAccountByIdController(usecases.MakeLoadAccountById())
}
