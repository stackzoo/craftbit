package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Fees represents information about recommended transaction fees.
type Fees struct {
	FastestFee  int64 `json:"fastestFee"`
	HalfHourFee int64 `json:"halfHourFee"`
	HourFee     int64 `json:"hourFee"`
	EconomyFee  int64 `json:"economyFee"`
	MinimumFee  int64 `json:"minimumFee"`
}

// GetFees retrieves the recommended transaction fees from the mempool.space API.
func GetFees() (Fees, error) {
	url := "https://mempool.space/api/v1/fees/recommended"
	resp, err := http.Get(url)
	if err != nil {
		return Fees{}, fmt.Errorf("failed to fetch fees: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Fees{}, fmt.Errorf("failed to fetch fees: status code %d", resp.StatusCode)
	}

	var fees Fees
	err = json.NewDecoder(resp.Body).Decode(&fees)
	if err != nil {
		return Fees{}, fmt.Errorf("failed to decode fees response: %v", err)
	}

	return fees, nil
}
