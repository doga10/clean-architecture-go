package account

import (
	"context"
	add "github.com/doga10/clean-architecture-go/src/data/protocols/db/account/add-account"
	load "github.com/doga10/clean-architecture-go/src/data/protocols/db/account/load-account-by-email"
	"github.com/doga10/clean-architecture-go/src/domain/models"
	"github.com/doga10/clean-architecture-go/src/domain/usecases/account"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AccountMongoRepository interface {
	add.AddAccountRepository
	load.LoadAccountByIdRepository
	account.LoadAccountById
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
	element, err := r.collection.InsertOne(context.TODO(), accountData)
	return element.InsertedID, err
}

func (r *repo) LoadByEmail(email string) (*models.AccountModel, error) {
	var element *models.AccountModel
	err := r.collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&element)
	if err != nil && err.Error() == "mongo: no documents in result" {
		return nil, nil
	}
	return element, err
}

func (r *repo) LoadById(id string) (*models.AccountModel, error) {
	var element *models.AccountModel
	parser, _ := primitive.ObjectIDFromHex(id)
	err := r.collection.FindOne(context.TODO(), bson.M{"_id": parser}).Decode(&element)
	if err != nil && err.Error() == "mongo: no documents in result" {
		return nil, nil
	}
	return element, err
}
