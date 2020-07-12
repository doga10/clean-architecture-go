package load_account_by_id

import (
	"github.com/doga10/clean-architecture-go/src/domain/usecases/account"
	"github.com/doga10/clean-architecture-go/src/presentation/protocols"
)

type LoadAccountByIdController interface {
	protocols.Controller
}

type controller struct {
	account.LoadAccountById
}

func NewLoadAccountByIdController(load account.LoadAccountById) LoadAccountByIdController {
	return &controller{load}
}

func (c *controller) Handle(req protocols.Request) protocols.Response {
	load, err := c.LoadById(req.Params.ByName("id"))
	if err != nil {
		return protocols.Response{500, err.Error()}
	}
	return protocols.Response{200, load}
}