package controllers

import (
	"math/rand"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RandomString(c echo.Context) error {
	var str = []string{"Hello", "Kamusta", "Xin chào", "Bonjour", "你好", "こんにちは", "안녕하십니까", "Hallo", "Merhaba", "Ciao", "Hola", "Привет", "Olá", "नमस्ते"}
	return c.String(http.StatusOK, str[rand.Intn(13)])

}
