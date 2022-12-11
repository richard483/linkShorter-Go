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

	e.Logger.Fatal(e.Start(":1323"))
}

func index(c echo.Context) error {
	return controllers.RandomString(c)
}

func getDatabaseList(c echo.Context) error {

	var database = configs.GetDatabaseList(c)

	return database
}

func getShortLinkCollectionList(c echo.Context) error {
	return controllers.GetShortLinkCollection(c)
}
