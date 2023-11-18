package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	usr      = "Admin"
	pwd      = "PASSWd"
	host     = "localhost"
	port     = 27017
	database = "test"
)

func GetCollection(collection string) *mongo.Collection {
	ctx := context.Background()
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d", usr, pwd, host, port)
	clientOption := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOption)
	if err != nil {
		panic(err.Error())
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err.Error())
	}

	return client.Database(database).Collection(collection)
}
