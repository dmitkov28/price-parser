package data

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type ArdesResults struct {
	Items []struct {
		ID             string `json:"id"`
		Prefix         string `json:"prefix"`
		Title          string `json:"title"`
		Image          string `json:"image"`
		Price          string `json:"price"`
		PricePromotion string `json:"price_promotion"`
		OnlinePrice    any    `json:"online_price"`
		Subcat         string `json:"subcat"`
		URL            string `json:"url"`
		Nomenclature   []any  `json:"nomenclature"`
		IsPromo        bool   `json:"isPromo"`
	} `json:"items"`
}

func FetchArdes(product string) ([]byte, error) {
	const ardesUrl = "https://ardes.bg/ajax?do=productAutocomplete"
	url := ardesUrl + fmt.Sprintf("&term=%s", product)
	data, err := Fetch(url)
	if err != nil {
		return nil, err
	}

	return data, nil
}

type ArdesAdapter struct{}

func (a *ArdesAdapter) Transform(rawData []byte) ([]Product, error) {
	var ardesData ArdesResults
	err := json.Unmarshal(rawData, &ardesData)
	if err != nil {
		return nil, err
	}

	var products []Product
	for _, item := range ardesData.Items {
		price, err := strconv.ParseFloat(item.Price, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse price for item %s: %w", item.ID, err)
		}

		product := Product{
			Name:  item.Title,
			Price: price,
			URL:   fmt.Sprintf("%s%s", "https://ardes.bg", item.URL),
		}
		products = append(products, product)
	}

	return products, nil
}
