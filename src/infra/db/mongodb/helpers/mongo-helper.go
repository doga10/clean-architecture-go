package helpers

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var (
	Url string
	DbName string
	Db *mongo.Database
)

func Connect(url string, name string) error {
	Url = url
	DbName = name
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		return err
	}
	Db = client.Database(name)
	return nil
}

func GetCollection(name string) *mongo.Collection {
	return Db.Collection(name)
}
