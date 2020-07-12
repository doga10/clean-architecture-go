package account

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/doga10/clean-architecture-go/src/domain/usecases/account"
	"github.com/doga10/clean-architecture-go/src/infra/db/mongodb/helpers"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
)

func Connect() *mongo.Collection {
	err := helpers.Connect("mongodb://127.0.0.1", "app")
	if err != nil {
		fmt.Println(err)
	}
	return helpers.GetCollection("accounts")
}

func TestNewAccountMongoRepository(t *testing.T) {
	collection := Connect()
	repo := NewAccountMongoRepository(collection)
	assert.NotNil(t, repo)
}

func TestRepo_Add(t *testing.T) {
	collection := Connect()
	repo := NewAccountMongoRepository(collection)

	var add account.AddAccountParams
	err := faker.FakeData(&add)
	if err != nil {
		fmt.Println(err)
	}
	cur, err := repo.Add(&add)
	assert.NotNil(t, cur)
	assert.Nil(t, err)
}

func TestRepo_LoadByEmail(t *testing.T) {
	collection := Connect()
	repo := NewAccountMongoRepository(collection)

	var add account.AddAccountParams
	err := faker.FakeData(&add)
	if err != nil {
		fmt.Println(err)
	}

	_, err = repo.Add(&add)
	assert.Nil(t, err)

	cur, err := repo.LoadByEmail(add.Email)
	assert.NotNil(t, cur)
	assert.Nil(t, err)

	cur2, err := repo.LoadByEmail(add.Name)
	assert.Nil(t, cur2)
	assert.Nil(t, err)
}
