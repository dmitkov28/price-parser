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

		rawArdesData, err := data.FetchArdes(encodedQuery)
		if err != nil {
			fmt.Printf("%v", err)
		}

		rawTechnoMarketData, err := data.FetchTechnoMarket(encodedQuery)

		if err != nil {
			fmt.Printf("%v", err)
		}

		ardesAdapter := data.ArdesAdapter{}
		ardesProducts, err := ardesAdapter.Transform(rawArdesData)
		if err != nil {
			fmt.Printf("%v", err)
		}

		technoMarketAdapter := data.TechnoMarketAdapter{}
		technoMarketProducts, err := technoMarketAdapter.Transform(rawTechnoMarketData)

		fmt.Println(technoMarketProducts)
		return render(c, templates.ProductComparison(ardesProducts, technoMarketProducts))
	}
}
