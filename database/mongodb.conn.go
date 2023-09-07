package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	user     = "alejandro"
	pwd      = ""
	host     = "localhost"
	port     = 27017
	database = "plantillas_bd_local"
)

func GetCollection(collection string) mongo.Collection {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d/?authSource=admin", user, pwd, host, port)

	fmt.Println("URI:", uri)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Cliente: ", client.Database(database))
	return *client.Database(database).Collection(collection)
}
