package main

import (
	"net/http"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", HelloWorld)
	e.Logger.Fatal(e.Start(":8080"))
}

func HelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!")
}
