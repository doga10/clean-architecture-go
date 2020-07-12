package account

import "github.com/doga10/clean-architecture-go/src/domain/models"

type LoadAccountById interface {
	LoadById(id string) (*models.AccountModel, error)
}
