package handler

import (
	"fmt"
	"net/url"

	"github.com/dmitkov28/price-parser/internal/data"
	"github.com/dmitkov28/price-parser/templates"
	"github.com/labstack/echo/v4"
)

func RootGETHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		return render(c, templates.Root())
	}
}

func SearchGetHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		query := c.QueryParam("query")
		encodedQuery := url.QueryEscape(query)

		rawData, err := data.FetchArdes(encodedQuery)
		if err != nil {
			fmt.Printf("%v", err)
		}
		ardesAdapter := data.ArdesAdapter{}
		products, err := ardesAdapter.Transform(rawData)
		if err != nil {
			fmt.Printf("%v", err)
		}
		return render(c, templates.ProductList(products))
	}
}
