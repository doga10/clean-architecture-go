package account

import (
	"context"
	add "github.com/doga10/clean-architecture-go/src/data/protocols/db/account/add-account"
	"github.com/doga10/clean-architecture-go/src/domain/usecases/account"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type AccountMongoRepository interface {
	add.AddAccountRepository
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
