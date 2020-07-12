package load_account_by_id

import (
	"fmt"
	load_account_by_id2 "github.com/doga10/clean-architecture-go/src/data/usecases/account/load-account-by-id"
	account2 "github.com/doga10/clean-architecture-go/src/infra/db/mongodb/account"
	"github.com/doga10/clean-architecture-go/src/infra/db/mongodb/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

type AddAccountParams struct {
	Name     string `bson:"name" json:"name" validate:"required" fake:"name"`
	Email    string `bson:"email" json:"email" validate:"required,email" fake:"email"`
	Password string `bson:"password" json:"password" validate:"required" fake:"name"`
}

func LoadAccountByIdControllerSpy() LoadAccountByIdController {
	err := helpers.Connect("mongodb://127.0.0.1:27017", "app")
	if err != nil {
		fmt.Sprintf("error connect mongodb %v", err)
	}
	collection := helpers.GetCollection("accounts")
	repo := account2.NewAccountMongoRepository(collection)
	data := load_account_by_id2.NewDbLoadAccountById(repo)

	return NewLoadAccountByIdController(data)
}

func TestNewLoadAccountByIdController(t *testing.T) {
	controller := LoadAccountByIdControllerSpy()
	assert.NotNil(t, controller)
}

//func TestController_Handle(t *testing.T) {
//	var req protocols.Request
//	controller := LoadAccountByIdControllerSpy()
//
//	res := controller.Handle(req)
//	assert.NotNil(t, res)
//	assert.Equal(t, res.Code, 200)
//}
