package main

import (
	"LinkShorter/configs"
	"LinkShorter/models"
	"LinkShorter/responses"
	"LinkShorter/src"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {

	e := echo.New()
	e.GET("/", index)
	e.GET("/databaseList", getDatabaseList)
	e.GET("/collectionList", getCollectionList)

	e.Logger.Fatal(e.Start(":1323"))
}

func index(c echo.Context) error {
	return c.String(http.StatusOK, src.RandomString())
}

func getDatabaseList(c echo.Context) error {

	var database = configs.GetDatabaseList(configs.DB)

	fmt.Println(database)

	return nil
}

func getCollectionList(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var shortLinks []models.ShortLink

	defer cancel()

	collection := configs.GetCollection(configs.DB, "shortLinks")

	res, err := collection.Find(ctx, bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ShortLinkResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
	}

	defer res.Close(ctx)

	for res.Next(ctx) {
		var shortLink models.ShortLink

		if err = res.Decode(&shortLink); err != nil {
			c.JSON(http.StatusInternalServerError, responses.ShortLinkResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		}

		shortLinks = append(shortLinks, shortLink)
	}

	return c.JSON(http.StatusOK, responses.ShortLinkResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": shortLinks}})

}
