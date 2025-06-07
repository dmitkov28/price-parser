package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/dmitkov28/price-parser/templates"
)

func RootGETHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		return render(c, templates.Root())
	}
}
