package validator

import (
	"github.com/doga10/clean-architecture-go/src/presentation/protocols"
	"github.com/go-playground/validator/v10"
)

type StructValidatorAdapter interface {
	protocols.Validation
}

type validators struct {
}

func NewStructValidatorAdapter() StructValidatorAdapter {
	return &validators{}
}

func (v *validators) Validate(u interface{}) error {
	validate := validator.New()
	err := validate.Struct(u)
	if err != nil {
		return err
	}
	return nil
}
