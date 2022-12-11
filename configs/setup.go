package configs

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))

	if err != nil {
		log.Fatal(err)

	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	return client
}

var DB *mongo.Client = connectDB()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("nephrenWeb").Collection(collectionName)
	return collection
}

func GetDatabaseList(client *mongo.Client) []string {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	database, err := client.ListDatabaseNames(ctx, bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	return database
}
