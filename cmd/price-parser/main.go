package main

import (
	"github.com/dmitkov28/price-parser/internal/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Static("/static", "assets")
	e.GET("/", handler.RootGETHandler())
	e.GET("/search", handler.SearchGetHandler())
	e.Logger.Fatal(e.Start(":1323"))
}
