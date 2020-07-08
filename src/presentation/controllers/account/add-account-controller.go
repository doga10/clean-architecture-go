package account

import (
	"encoding/json"
	"github.com/doga10/clean-architecture-go/src/domain/usecases/account"
	"github.com/doga10/clean-architecture-go/src/presentation/protocols"
	"io/ioutil"
)

type AddAccountController interface {
	protocols.Controller
}

type controller struct {
	account.AddAccount
}

func NewAddAccountController(add account.AddAccount) AddAccountController {
	return &controller{
		add,
	}
}

func (c *controller) Handle(req protocols.Request) protocols.Response {
	var addAccountParams *account.AddAccountParams
	parser, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return protocols.Response{500, err.Error()}
	}

	err = json.Unmarshal(parser, &addAccountParams)
	if err != nil {
		return protocols.Response{500, err.Error()}
	}

	addAccount, err := c.Add(addAccountParams)
	if err != nil {
		return protocols.Response{500, err.Error()}
	}
	return protocols.Response{201, addAccount}
}
