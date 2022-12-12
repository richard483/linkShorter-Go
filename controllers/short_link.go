package controllers

import (
	"LinkShorter/models"
	"LinkShorter/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GoToPage(c echo.Context) error {

	shortLink := new(models.ShortLink)

	if err := c.Bind(&shortLink); err != nil {
		return c.JSON(http.StatusBadRequest, responses.ShortLinkResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
	}

	shortLink = GetShortLinkbyName(c, shortLink.Name)

	if shortLink == nil {
		return c.JSON(http.StatusNotFound, responses.ShortLinkResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "not found"}})
	}

	var redirect = c.Redirect(http.StatusFound, shortLink.LongLink)
	return redirect
}
