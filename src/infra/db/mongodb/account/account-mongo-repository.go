package account

import (
	"context"
	add "github.com/doga10/clean-architecture-go/src/data/protocols/db/account/add-account"
	load "github.com/doga10/clean-architecture-go/src/data/protocols/db/account/load-account-by-email"
	"github.com/doga10/clean-architecture-go/src/domain/models"
	"github.com/doga10/clean-architecture-go/src/domain/usecases/account"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type AccountMongoRepository interface {
	add.AddAccountRepository
	load.LoadAccountByIdRepository
}

type repo struct {
	collection *mongo.Collection
}

func NewAccountMongoRepository(collection *mongo.Collection) AccountMongoRepository {
	return &repo{
		collection: collection,
	}
}

func (r *repo) Add(accountData *account.AddAccountParams) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	element, err := r.collection.InsertOne(ctx, accountData)
	if err != nil {
		return nil, err
	}

	return element, nil
}

func (r *repo) LoadByEmail(email string) (*models.AccountModel, error) {
	var element *models.AccountModel
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(&element)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return nil, nil
		}
		return nil, err
	}

	return element, nil
}
