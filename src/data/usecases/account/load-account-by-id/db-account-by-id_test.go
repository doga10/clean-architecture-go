package load_account_by_id

import (
	"context"
	"fmt"
	"github.com/doga10/clean-architecture-go/src/domain/usecases/account"
	account2 "github.com/doga10/clean-architecture-go/src/infra/db/mongodb/account"
	"github.com/doga10/clean-architecture-go/src/infra/db/mongodb/helpers"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

type MockAccount struct {
	Email              string  `faker:"email"`
}

var ID string

func LoadAccountByIdSpy() DbLoadAccountById {
	var add account.AddAccountParams
	add.Name = "devops"
	add.Email = "devops@devops.com"
	add.Password = "devops"
	err := helpers.Connect("mongodb://127.0.0.1", "app")
	if err != nil {
		fmt.Println(err)
	}
	collection := helpers.GetCollection("accounts")
	result, _ := collection.InsertOne(context.TODO(), add)
	ID = result.InsertedID.(primitive.ObjectID).Hex()
	repo := account2.NewAccountMongoRepository(collection)
	return NewDbLoadAccountById(repo)
}

func TestNewDbLoadAccountById(t *testing.T) {
	repo := LoadAccountByIdSpy()
	assert.NotNil(t, repo)
}

func TestSvc_LoadById(t *testing.T) {
	repo := LoadAccountByIdSpy()
	assert.NotNil(t, repo)

	cur, err := repo.LoadById(ID)
	assert.Nil(t, err)
	assert.NotNil(t, cur)
}

func TestSvc_LoadByIdNotFound(t *testing.T) {
	repo := LoadAccountByIdSpy()
	assert.NotNil(t, repo)

	cur, err := repo.LoadById("t20@test.com")
	assert.Nil(t, err)
	assert.Nil(t, cur)
}

