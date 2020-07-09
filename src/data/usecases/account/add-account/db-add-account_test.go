package add_account

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/doga10/clean-architecture-go/src/domain/usecases/account"
	crypto "github.com/doga10/clean-architecture-go/src/infra/criptography/bcrypt-adapter"
	account2 "github.com/doga10/clean-architecture-go/src/infra/db/mongodb/account"
	"github.com/doga10/clean-architecture-go/src/infra/db/mongodb/helpers"
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
	if test == nil {
		t.Error("error start module add account")
	}
}

func TestSvc_Add(t *testing.T) {
	test := AddAccountSpy()
	if test == nil {
		t.Error("error start module add account")
	}

	var add account.AddAccountParams
	err := faker.FakeData(&add)
	if err != nil {
		fmt.Println(err)
	}

	collection := helpers.GetCollection("accounts")
	repo := account2.NewAccountMongoRepository(collection)
	cur, err := repo.LoadByEmail(add.Email)
	if cur != nil {
		t.Error("error register user")
	}

	crypt := crypto.NewBcryptAdapter()
	pass, err := crypt.Hash(add.Password)
	if err != nil {
		t.Errorf("erro crypt pass %v", err)
	}
	add.Password = pass

	_, _ = test.Add(&add)
}

func TestSvc_AddError(t *testing.T) {
	test := AddAccountSpy()
	if test == nil {
		t.Error("error start module add account")
	}

	var add account.AddAccountParams
	err := faker.FakeData(&add)
	if err != nil {
		fmt.Println(err)
	}

	collection := helpers.GetCollection("accounts")
	repo := account2.NewAccountMongoRepository(collection)
	cur, err := repo.LoadByEmail(add.Email)
	if cur != nil {
		t.Error("not found user")
	}
	if err != nil {
		t.Errorf(err.Error())
	}
}
