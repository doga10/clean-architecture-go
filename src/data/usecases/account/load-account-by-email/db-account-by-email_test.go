package load_account_by_email

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	account2 "github.com/doga10/clean-architecture-go/src/infra/db/mongodb/account"
	"github.com/doga10/clean-architecture-go/src/infra/db/mongodb/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockAccount struct {
	Email              string  `faker:"email"`
}

func LoadAccountByEmailSpy() DbLoadAccountByEmail {
	err := helpers.Connect("mongodb://127.0.0.1", "app")
	if err != nil {
		fmt.Println(err)
	}
	collection := helpers.GetCollection("accounts")
	repo := account2.NewAccountMongoRepository(collection)
	return NewDbLoadAccountByEmail(repo)
}

func TestNewDbLoadAccountByEmail(t *testing.T) {
	repo := LoadAccountByEmailSpy()
	assert.NotNil(t, repo)
}

func TestSvc_LoadByEmail(t *testing.T) {
	repo := LoadAccountByEmailSpy()
	assert.NotNil(t, repo)

	cur, err := repo.LoadByEmail("t20@test.com")
	assert.Nil(t, err)
	assert.NotNil(t, cur)
}

func TestSvc_LoadByEmailNotFound(t *testing.T) {
	mock := MockAccount{}
	err := faker.FakeData(&mock)
	if err != nil {
		fmt.Println(err)
	}
	repo := LoadAccountByEmailSpy()
	if repo == nil {
		t.Error("error start module add account")
	}

	cur, err := repo.LoadByEmail(mock.Email)
	assert.Nil(t, err)
	assert.Nil(t, cur)
}

