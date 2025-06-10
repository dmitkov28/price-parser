package data

import (
	"encoding/json"
	"fmt"
)

type TechnoMarketResults struct {
	Suggestions []interface{} `json:"suggestions"`
	Results     []struct {
		Title      string   `json:"title"`
		Type       string   `json:"type"`
		Keyword    string   `json:"keyword,omitempty"`
		URL        string   `json:"url"`
		Image      string   `json:"image"`
		Brand      string   `json:"brand,omitempty"`
		Categories []string `json:"categories,omitempty"`
		Code       string   `json:"code,omitempty"`
		Price      int      `json:"price,omitempty"`
	} `json:"results"`
	Categories []struct {
		Title string `json:"title"`
		Count int    `json:"count"`
		Slug  string `json:"slug"`
	} `json:"categories"`
}

func FetchTechnoMarket(product string) ([]byte, error) {
	const technomarketUrl = "https://www.technomarket.bg/api/frontend/search"
	url := technomarketUrl + fmt.Sprintf("?query=%s", product)
	data, err := Fetch(url)
	if err != nil {
		return nil, err
	}

	return data, nil
}

type TechnoMarketAdapter struct{}

func (a *TechnoMarketAdapter) Transform(rawData []byte) ([]Product, error) {
	var technomarketData TechnoMarketResults
	err := json.Unmarshal(rawData, &technomarketData)
	if err != nil {
		return nil, err
	}

	var products []Product
	for _, item := range technomarketData.Results {
		price := item.Price

		product := Product{
			Name:  item.Title,
			Price: float64(price),
			URL:   fmt.Sprintf("%s%s", "https://technomarket.bg", item.URL),
		}
		products = append(products, product)
	}

	return products, nil
}
