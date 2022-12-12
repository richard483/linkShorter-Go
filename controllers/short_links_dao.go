package controllers

import (
	"LinkShorter/configs"
	"LinkShorter/models"
	"LinkShorter/responses"
	"context"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var shortLinkCollection *mongo.Collection = configs.GetCollection(configs.DB, "shortLinks")
var validate = validator.New()

func GetShortLinkCollection(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	var shortLinks []models.ShortLink

	defer cancel()

	result, err := shortLinkCollection.Find(ctx, bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ShortLinkResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
	}

	defer result.Close(ctx)

	for result.Next(ctx) {
		var shortLink models.ShortLink

		if err = result.Decode(&shortLink); err != nil {
			c.JSON(http.StatusInternalServerError, responses.ShortLinkResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		}

		shortLinks = append(shortLinks, shortLink)
	}

	return c.JSON(http.StatusOK, responses.ShortLinkResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": shortLinks}})
}

func GetShortLinkbyName(c echo.Context, name string) *models.ShortLink {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var shortLink = new(models.ShortLink)

	defer cancel()

	err := shortLinkCollection.FindOne(ctx, bson.M{"name": name}).Decode(&shortLink)

	if err != nil {
		return nil
	}
	return shortLink
}

func GetShortLinkDatabyName(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	shortLink := new(models.ShortLink)

	defer cancel()

	if err := c.Bind(&shortLink); err != nil {
		return c.JSON(http.StatusBadRequest, responses.ShortLinkResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
	}

	defer cancel()

	err := shortLinkCollection.FindOne(ctx, bson.M{"name": shortLink.Name}).Decode(&shortLink)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.ShortLinkResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
	}

	return c.JSON(http.StatusOK, responses.ShortLinkResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": shortLink}})

}

func CreateShortLinkCollection(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	shortLink := new(models.ShortLink)

	defer cancel()

	if err := c.Bind(&shortLink); err != nil {
		return c.JSON(http.StatusBadRequest, responses.ShortLinkResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
	}

	if validateError := validate.Struct(shortLink); validateError != nil {
		return c.JSON(http.StatusBadRequest, responses.ShortLinkResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validateError.Error()}})
	}

	err := shortLinkCollection.FindOne(ctx, bson.M{"name": shortLink.Name}).Decode(&shortLink)

	if err == nil {
		return c.JSON(http.StatusInternalServerError, responses.ShortLinkResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "found duplicate name"}})
	}

	newShortLink := models.ShortLink{
		Name:     shortLink.Name,
		LongLink: shortLink.LongLink,
	}

	result, err := shortLinkCollection.InsertOne(ctx, newShortLink)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.ShortLinkResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
	}

	return c.JSON(http.StatusOK, responses.ShortLinkResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": result}})
}
