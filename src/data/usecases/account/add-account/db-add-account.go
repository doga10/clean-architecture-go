package add_account

import (
	"errors"
	add "github.com/doga10/clean-architecture-go/src/data/protocols/db/account/add-account"
	load "github.com/doga10/clean-architecture-go/src/data/protocols/db/account/load-account-by-email"
	"github.com/doga10/clean-architecture-go/src/domain/usecases/account"
	bcrypt_adapter "github.com/doga10/clean-architecture-go/src/infra/criptography/bcrypt-adapter"
)

type DbAddAccount interface {
	account.AddAccount
}

type svc struct {
	add add.AddAccountRepository
	load load.LoadAccountByIdRepository
	crypto bcrypt_adapter.BcryptAdapter
}

func NewDbAddAccount(function add.AddAccountRepository, load load.LoadAccountByIdRepository, crypto bcrypt_adapter.BcryptAdapter) DbAddAccount {
	return &svc{
		function,
		load,
		crypto,
	}
}

func (s *svc) Add(accountData *account.AddAccountParams) (interface{}, error) {
	element, err := s.load.LoadByEmail(accountData.Email)
	if err != nil {
		return nil, err
	}
	if element == nil {
		password, err := s.crypto.Hash(accountData.Password)
		accountData.Password = password
		element, err := s.add.Add(accountData)
		if err != nil {
			return nil, err
		}

		return element, nil
	}

	return nil, errors.New("E-mail já registrado")
}
