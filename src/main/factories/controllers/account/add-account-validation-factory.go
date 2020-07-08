package account

import (
	"github.com/doga10/clean-architecture-go/src/infra/validator"
	"github.com/doga10/clean-architecture-go/src/presentation/protocols"
)

func MakeAddAccountValidation() protocols.Validation {
	return validator.NewStructValidatorAdapter()
}

//export const makeSignUpValidation = (): ValidationComposite => {
//  const validations: Validation[] = []
//  for (const field of ['name', 'email', 'password', 'passwordConfirmation']) {
//    validations.push(new RequiredFieldValidation(field))
//  }
//  validations.push(new CompareFieldsValidation('password', 'passwordConfirmation'))
//  validations.push(new EmailValidation('email', new EmailValidatorAdapter()))
//  return new ValidationComposite(validations)
//}