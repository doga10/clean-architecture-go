package load_account_by_email

import (
	load "github.com/doga10/clean-architecture-go/src/data/protocols/db/account/load-account-by-email"
	"github.com/doga10/clean-architecture-go/src/domain/models"
	"github.com/doga10/clean-architecture-go/src/domain/usecases/account"
)

type DbLoadAccountByEmail interface {
	account.LoadAccountByEmail
}

type svc struct {
	loadByEmail load.LoadAccountByIdRepository
}

func NewDbLoadAccountByEmail(function load.LoadAccountByIdRepository) DbLoadAccountByEmail {
	return &svc{
		loadByEmail: function,
	}
}

func (s *svc) LoadByEmail(email string) (*models.AccountModel, error) {
	element, err := s.loadByEmail.LoadByEmail(email)
	if err != nil {
		return nil, err
	}

	return element, nil
}
