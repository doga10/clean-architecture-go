package add_account

import (
	"encoding/json"
	"fmt"
	"github.com/bxcodec/faker/v3"
	add_account "github.com/doga10/clean-architecture-go/src/data/usecases/account/add-account"
	load_account_by_email "github.com/doga10/clean-architecture-go/src/data/usecases/account/load-account-by-email"
	bcrypt_adapter "github.com/doga10/clean-architecture-go/src/infra/criptography/bcrypt-adapter"
	account2 "github.com/doga10/clean-architecture-go/src/infra/db/mongodb/account"
	"github.com/doga10/clean-architecture-go/src/infra/db/mongodb/helpers"
	"github.com/doga10/clean-architecture-go/src/infra/validator"
	"github.com/doga10/clean-architecture-go/src/presentation/protocols"
	"github.com/stretchr/testify/assert"
	"testing"
)

type AddAccountParams struct {
	Name     string `bson:"name" json:"name" validate:"required" fake:"name"`
	Email    string `bson:"email" json:"email" validate:"required,email" fake:"email"`
	Password string `bson:"password" json:"password" validate:"required" fake:"name"`
}

func AddAccountControllerSpy() AddAccountController {
	err := helpers.Connect("mongodb://127.0.0.1:27017", "app")
	if err != nil {
		fmt.Sprintf("error connect mongodb %v", err)
	}
	crypt := bcrypt_adapter.NewBcryptAdapter()
	collection := helpers.GetCollection("accounts")
	repo := account2.NewAccountMongoRepository(collection)
	load_account_by_email.NewDbLoadAccountByEmail(repo)
	svc := add_account.NewDbAddAccount(repo, repo, crypt)
	valid := validator.NewStructValidatorAdapter()

	return NewAddAccountController(svc, valid)
}

func TestNewAddAccountController(t *testing.T) {
	controller := AddAccountControllerSpy()
	assert.NotNil(t, controller)
}

func TestController_Handle(t *testing.T) {
	var req protocols.Request
	var body AddAccountParams
	err := faker.FakeData(&body)
	if err != nil {
		fmt.Println(err)
	}
	body.Email = body.Email + "@gmail.com"
	controller := AddAccountControllerSpy()
	b, err := json.Marshal(body)
	assert.Nil(t, err)
	req.Body = b

	res := controller.Handle(req)
	assert.NotNil(t, res)
	assert.Equal(t, res.Code, 201)
}

func TestController_HandleNotValid(t *testing.T) {
	var req protocols.Request
	var body AddAccountParams
	err := faker.FakeData(&body)
	if err != nil {
		fmt.Println(err)
	}
	controller := AddAccountControllerSpy()
	b, err := json.Marshal(body)
	assert.Nil(t, err)
	req.Body = b

	res := controller.Handle(req)
	assert.NotNil(t, res)
	assert.Equal(t, res.Code, 400)
}