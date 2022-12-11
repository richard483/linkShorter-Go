package main

import (
	"LinkShorter/configs"
	"LinkShorter/controllers"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	configs.ConnectDB()

	e.GET("/", index)
	e.GET("/databaseList", getDatabaseList)
	e.GET("/collectionList", getShortLinkCollectionList)
	e.POST("/shortLink", insertShortLink)
	e.POST("/findLink", findShortLink)

	e.Logger.Fatal(e.Start(":1323"))
}

func index(c echo.Context) error {
	var randomhello = controllers.RandomString(c)
	return randomhello
}

func getDatabaseList(c echo.Context) error {

	var database = configs.GetDatabaseList(c)

	return database
}

func getShortLinkCollectionList(c echo.Context) error {
	var getShortLinkCollection = controllers.GetShortLinkCollection(c)
	return getShortLinkCollection
}

func insertShortLink(c echo.Context) error {
	var createShortLink = controllers.CreateShortLinkCollection(c)
	return createShortLink
}

func findShortLink(c echo.Context) error {
	var findShortLink = controllers.GetShortLinkbyName(c)
	return findShortLink
}
