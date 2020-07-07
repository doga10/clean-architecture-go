package add_account

import (
	add "github.com/doga10/clean-architecture-go/src/data/protocols/db/account/add-account"
	"github.com/doga10/clean-architecture-go/src/domain/usecases/account"
)

type DbAddAccount interface {
	account.AddAccount
}

type svc struct {
	add add.AddAccountRepository
}

func NewDbAddAccount(function add.AddAccountRepository) DbAddAccount {
	return &svc{
		add: function,
	}
}

func (s *svc) Add(accountData *account.AddAccountParams) (interface{}, error) {
	element, err := s.add.Add(accountData)
	if err != nil {
		return nil, err
	}

	return element, nil
}
