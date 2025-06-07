package main

import (
	"github.com/dmitkov28/price-parser/internal/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", handler.RootGETHandler())
	e.Logger.Fatal(e.Start(":1323"))
}
