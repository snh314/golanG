package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var Client *mongo.Client = DBSet()

func DBSet() *mongo.Client {
	client, err := mongo.newClient(options.Client().ApplyURI("mongodb: //localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Panicln("failed to connect to mongodb :( ")
		return nil

	}
	fmt.Println("Successfully connect to mongodb ")
	return client

}

func UserData(client *mongo.Client, collectionName string) *mongo.Collection {

}
func ProductData(client *mongo.Client, collection string) *mongo.Collection {

}
