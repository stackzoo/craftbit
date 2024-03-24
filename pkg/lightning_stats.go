package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// LightningStatistics represents the latest lightning network statistics.
type LightningStatistics struct {
	Latest   LightningInfo `json:"latest"`
	Previous LightningInfo `json:"previous"`
}

// LightningInfo represents information about lightning network statistics.
type LightningInfo struct {
	ChannelCount       int   `json:"channel_count"`
	NodeCount          int   `json:"node_count"`
	TotalCapacity      int64 `json:"total_capacity"`
	TorNodes           int   `json:"tor_nodes"`
	ClearnetNodes      int   `json:"clearnet_nodes"`
	UnannouncedNodes   int   `json:"unannounced_nodes"`
	AverageCapacity    int64 `json:"avg_capacity"`
	AverageFeeRate     int   `json:"avg_fee_rate"`
	AverageBaseFeeMtok int   `json:"avg_base_fee_mtokens"`
	MedianCapacity     int64 `json:"med_capacity"`
	MedianFeeRate      int   `json:"med_fee_rate"`
	MedianBaseFeeMtok  int   `json:"med_base_fee_mtokens"`
	ClearnetTorNodes   int   `json:"clearnet_tor_nodes"`
}

// GetLightningStatistics retrieves the latest lightning network statistics from the mempool.space API.
func GetLightningStatistics() (LightningStatistics, error) {
	url := "https://mempool.space/api/v1/lightning/statistics/latest"
	resp, err := http.Get(url)
	if err != nil {
		return LightningStatistics{}, fmt.Errorf("failed to fetch lightning statistics: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return LightningStatistics{}, fmt.Errorf("failed to fetch lightning statistics: status code %d", resp.StatusCode)
	}

	var lightningStats LightningStatistics
	err = json.NewDecoder(resp.Body).Decode(&lightningStats)
	if err != nil {
		return LightningStatistics{}, fmt.Errorf("failed to decode lightning statistics response: %v", err)
	}

	return lightningStats, nil
}
