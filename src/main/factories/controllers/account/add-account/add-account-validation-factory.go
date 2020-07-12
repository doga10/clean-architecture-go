package add_account

import (
	"github.com/doga10/clean-architecture-go/src/infra/validator"
	"github.com/doga10/clean-architecture-go/src/presentation/protocols"
)

func MakeAddAccountValidation() protocols.Validation {
	return validator.NewStructValidatorAdapter()
}
