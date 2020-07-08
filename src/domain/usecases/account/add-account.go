package account

import (
	"github.com/go-playground/validator/v10"
)

type AddAccountParams struct {
	Name     string `bson:"name" json:"name" validate:"required"`
	Email    string `bson:"email" json:"email" validate:"required,email"`
	Password string `bson:"password" json:"password" validate:"required"`
}

type AddAccount interface {
	Add(accountData *AddAccountParams) (interface{}, error)
}

func (u AddAccountParams) Validate() error {
	validate := validator.New()
	err := validate.Struct(u)
	if err != nil {
		return err
	}
	return nil
}
