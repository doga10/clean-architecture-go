package load_account_by_email

import (
	"github.com/doga10/clean-architecture-go/src/domain/models"
)

type LoadAccountByIdRepository interface {
	LoadByEmail(email string) (*models.AccountModel, error)
}
