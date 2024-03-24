package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strings"
)

// LightningNode represents information about a lightning network node.
type LightningNode struct {
	PublicKey string `json:"publicKey"`
	Alias     string `json:"alias"`
	Capacity  int64  `json:"capacity,omitempty"`
	Channels  int    `json:"channels,omitempty"`
}

// LightningTopNodes represents the top nodes by liquidity and connectivity.
type LightningTopNodes struct {
	TopByCapacity []LightningNode `json:"topByCapacity"`
	TopByChannels []LightningNode `json:"topByChannels"`
}

// GetLightningTopNodes retrieves the top nodes by liquidity and connectivity from the mempool.space API.
func GetLightningTopNodes() (LightningTopNodes, error) {
	url := "https://mempool.space/api/v1/lightning/nodes/rankings"
	resp, err := http.Get(url)
	if err != nil {
		return LightningTopNodes{}, fmt.Errorf("failed to fetch top lightning nodes: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return LightningTopNodes{}, fmt.Errorf("failed to fetch top lightning nodes: status code %d", resp.StatusCode)
	}

	var topNodes LightningTopNodes
	err = json.NewDecoder(resp.Body).Decode(&topNodes)
	if err != nil {
		return LightningTopNodes{}, fmt.Errorf("failed to decode top lightning nodes response: %v", err)
	}

	return topNodes, nil
}

// ByChannelsString returns a string representation of top nodes sorted by channels.
func (tn LightningTopNodes) ByChannelsString() string {
	sort.Slice(tn.TopByChannels, func(i, j int) bool {
		return tn.TopByChannels[i].Channels > tn.TopByChannels[j].Channels
	})

	var builder strings.Builder
	for _, node := range tn.TopByChannels {
		builder.WriteString(fmt.Sprintf("Alias: %s, Channels: %d\n", node.Alias, node.Channels))
	}
	return builder.String()
}

// ByCapacityString returns a string representation of top nodes sorted by capacity.
func (tn LightningTopNodes) ByCapacityString() string {
	sort.Slice(tn.TopByCapacity, func(i, j int) bool {
		return tn.TopByCapacity[i].Capacity > tn.TopByCapacity[j].Capacity
	})

	var builder strings.Builder
	for _, node := range tn.TopByCapacity {
		builder.WriteString(fmt.Sprintf("Alias: %s, Capacity: %d\n", node.Alias, node.Capacity))
	}
	return builder.String()
}
