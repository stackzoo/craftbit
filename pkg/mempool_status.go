package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// MempoolStatus represents the DTO (Data Transfer Object) for mempool status.
type MempoolStatus struct {
	Count    int `json:"count"`
	VSize    int `json:"vsize"`
	TotalFee int `json:"total_fee"`
	// FeeHistogram []struct{} `json:"-"`
}

// GetMempoolStatus retrieves mempool status from the mempool.space API.
func GetMempoolStatus() (MempoolStatus, error) {
	url := "https://mempool.space/api/mempool"
	resp, err := http.Get(url)
	if err != nil {
		return MempoolStatus{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return MempoolStatus{}, fmt.Errorf("failed to fetch mempool status: status code %d", resp.StatusCode)
	}

	var mempoolStatus MempoolStatus
	err = json.NewDecoder(resp.Body).Decode(&mempoolStatus)
	if err != nil {
		return MempoolStatus{}, err
	}

	// Remove the fee_histogram property
	// mempoolStatus.FeeHistogram = nil

	return mempoolStatus, nil
}
