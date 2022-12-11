package controllers

import (
	"LinkShorter/configs"
	"LinkShorter/models"
	"LinkShorter/responses"
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func GetShortLinkCollection(c echo.Context) error {
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
