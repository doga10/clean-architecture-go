package add_account

import (
	"context"
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/doga10/clean-architecture-go/src/domain/usecases/account"
	crypto "github.com/doga10/clean-architecture-go/src/infra/criptography/bcrypt-adapter"
	account2 "github.com/doga10/clean-architecture-go/src/infra/db/mongodb/account"
	"github.com/doga10/clean-architecture-go/src/infra/db/mongodb/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func AddAccountSpy() DbAddAccount {
	err := helpers.Connect("mongodb://127.0.0.1", "app")
	if err != nil {
		fmt.Println(err)
	}
	crypt := crypto.NewBcryptAdapter()
	collection := helpers.GetCollection("accounts")
	repo := account2.NewAccountMongoRepository(collection)
	return NewDbAddAccount(repo, repo, crypt)
}

func TestNewDbAddAccount(t *testing.T) {
	test := AddAccountSpy()
	assert.NotNil(t, test)
}

func TestSvc_Add(t *testing.T) {
	test := AddAccountSpy()
	assert.NotNil(t, test)

	var add account.AddAccountParams
	err := faker.FakeData(&add)
	if err != nil {
		fmt.Println(err)
	}

	crypt := crypto.NewBcryptAdapter()
	pass, err := crypt.Hash(add.Password)
	assert.Nil(t, err)
	add.Password = pass

	cur, err := test.Add(&add)
	assert.NotNil(t, cur)
	assert.Nil(t, err)

	cur2, err := test.Add(&add)
	assert.Nil(t, cur2)
	assert.NotNil(t, err)
}

func TestSvc_AddError(t *testing.T) {
	test := AddAccountSpy()
	assert.NotNil(t, test)

	var add account.AddAccountParams
	err := faker.FakeData(&add)
	if err != nil {
		fmt.Println(err)
	}

	crypt := crypto.NewBcryptAdapter()
	pass, err := crypt.Hash(add.Password)
	assert.Nil(t, err)
	add.Password = pass

	helpers.Client.Disconnect(context.TODO())

	cur, err := test.Add(&add)
	assert.Nil(t, cur)
	assert.NotNil(t, err)
}
