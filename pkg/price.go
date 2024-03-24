// pkg/price.go

package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Transaction represents information about a single transaction.
type Price struct {
	//Time int64 `json:"time"`
	USD int64 `json:"USD"`
	EUR int64 `json:"EUR"`
	GBP int64 `json:"GBP"`
	CAD int64 `json:"CAD"`
	CHF int64 `json:"CHF"`
	AUD int64 `json:"AUD"`
	JPY int64 `json:"JPY"`
}

// GetPrices retrieves the current prices from the mempool.space API.
func GetPrices() (Price, error) {
	url := "https://mempool.space/api/v1/prices"
	resp, err := http.Get(url)
	if err != nil {
		return Price{}, fmt.Errorf("failed to fetch prices: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Price{}, fmt.Errorf("failed to fetch prices: status code %d", resp.StatusCode)
	}

	var prices Price
	err = json.NewDecoder(resp.Body).Decode(&prices)
	if err != nil {
		return Price{}, fmt.Errorf("failed to decode prices response: %v", err)
	}

	return prices, nil
}
