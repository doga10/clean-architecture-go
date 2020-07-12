package load_account_by_id

import (
	load "github.com/doga10/clean-architecture-go/src/data/protocols/db/account/load-account-by-id"
	"github.com/doga10/clean-architecture-go/src/domain/models"
	"github.com/doga10/clean-architecture-go/src/domain/usecases/account"
)

type DbLoadAccountById interface {
	account.LoadAccountById
}

type svc struct {
	loadById load.LoadAccountByIdRepository
}

func NewDbLoadAccountById(function load.LoadAccountByIdRepository) DbLoadAccountById {
	return &svc{
		loadById: function,
	}
}

func (s *svc) LoadById(id string) (*models.AccountModel, error) {
	element, err := s.loadById.LoadById(id)
	return element, err
}
