package account

import (
	"fmt"
	"github.com/doga10/clean-architecture-go/src/domain/usecases/account"
	"github.com/doga10/clean-architecture-go/src/infra/db/mongodb/helpers"
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
	t.Parallel()
	collection := Connect()
	repo := NewAccountMongoRepository(collection)
	if repo == nil {
		t.Errorf("error class NewAccountMongoRepository %v", repo)
	}
}

func TestRepo_Add(t *testing.T) {
	t.Parallel()
	collection := Connect()
	repo := NewAccountMongoRepository(collection)
	if repo == nil {
		t.Errorf("error class NewAccountMongoRepository %v", repo)
	}
	var add account.AddAccountParams
	add.Name = "test"
	add.Email = "test@test.com"
	add.Password = "password"
	_, err := repo.Add(&add)
	if err != nil {
		t.Errorf("error add account %v", err)
	}
}

func TestRepo_LoadByEmail(t *testing.T) {
	t.Parallel()
	collection := Connect()
	repo := NewAccountMongoRepository(collection)
	if repo == nil {
		t.Errorf("error class NewAccountMongoRepository %v", repo)
	}
	var add account.AddAccountParams
	add.Name = "test"
	add.Email = "test@test.com"
	add.Password = "password"
	_, err := repo.Add(&add)
	if err != nil {
		t.Errorf("error add account %v", err)
	}

	_, err = repo.LoadByEmail(add.Email)
	if err != nil {
		t.Errorf("error load account by email %v", err)
	}
}

func TestRepo_LoadByEmailError(t *testing.T) {
	t.Parallel()
	collection := Connect()
	repo := NewAccountMongoRepository(collection)
	if repo == nil {
		t.Errorf("error class NewAccountMongoRepository %v", repo)
	}
	var add account.AddAccountParams
	add.Name = "test"
	add.Email = "test@test.com"
	add.Password = "password"
	_, err := repo.Add(&add)
	if err != nil {
		t.Errorf("error add account %v", err)
	}

	_, err = repo.LoadByEmail(add.Name)
	if err != nil {
		t.Errorf("error param load account by email  %v", err)
	}

	cur, err := repo.LoadByEmail("")
	fmt.Println(cur)
	fmt.Println(err)
}
