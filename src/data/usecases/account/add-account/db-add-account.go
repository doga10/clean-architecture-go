package add_account

import (
	"fmt"
	add "github.com/doga10/clean-architecture-go/src/data/protocols/db/account/add-account"
	"github.com/doga10/clean-architecture-go/src/domain/usecases/account"
	bcrypt_adapter "github.com/doga10/clean-architecture-go/src/infra/criptography/bcrypt-adapter"
)

type DbAddAccount interface {
	account.AddAccount
}

type svc struct {
	add add.AddAccountRepository
	crypto bcrypt_adapter.BcryptAdapter
}

func NewDbAddAccount(function add.AddAccountRepository, crypto bcrypt_adapter.BcryptAdapter) DbAddAccount {
	return &svc{
		add: function,
		crypto: crypto,
	}
}

func (s *svc) Add(accountData *account.AddAccountParams) (interface{}, error) {
	password, err := s.crypto.Hash(accountData.Password)
	accountData.Password = password
	element, err := s.add.Add(accountData)
	if err != nil {
		return nil, err
	}

	return element, nil
}
