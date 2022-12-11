package configs

import (
	"LinkShorter/responses"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
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

var DB *mongo.Client = ConnectDB()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("nephrenWeb").Collection(collectionName)
	return collection
}

func GetDatabaseList(c echo.Context) error {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	database, err := DB.ListDatabaseNames(ctx, bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, responses.ShortLinkResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": database}})
}
