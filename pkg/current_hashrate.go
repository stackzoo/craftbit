package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// CurrentHashrate represents the DTO (Data Transfer Object) for current hashrate data.
type CurrentHashrate struct {
	CurrentHashrate   float64 `json:"currentHashrate"`
	CurrentDifficulty float64 `json:"currentDifficulty"`
}

// GetCurrentHashrate retrieves current hashrate data from the mempool.space API.
func GetCurrentHashrate() (CurrentHashrate, error) {
	url := "https://mempool.space/api/v1/mining/hashrate/3d"
	resp, err := http.Get(url)
	if err != nil {
		return CurrentHashrate{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return CurrentHashrate{}, fmt.Errorf("failed to fetch current hashrate: status code %d", resp.StatusCode)
	}

	var hashrateData map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&hashrateData)
	if err != nil {
		return CurrentHashrate{}, err
	}

	// Extract only currentHashrate and currentDifficulty from the response
	currentHashrate, _ := hashrateData["currentHashrate"].(float64)
	currentDifficulty, _ := hashrateData["currentDifficulty"].(float64)

	return CurrentHashrate{CurrentHashrate: currentHashrate, CurrentDifficulty: currentDifficulty}, nil
}
