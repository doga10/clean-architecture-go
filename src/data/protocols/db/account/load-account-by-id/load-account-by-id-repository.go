package load_account_by_id

import (
	"github.com/doga10/clean-architecture-go/src/domain/models"
)

type LoadAccountByIdRepository interface {
	LoadById(id string) (*models.AccountModel, error)
}
