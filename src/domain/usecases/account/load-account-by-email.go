package account

import "github.com/doga10/clean-architecture-go/src/domain/models"

type LoadAccountByEmail interface {
	LoadByEmail(email string) (*models.AccountModel, error)
}
