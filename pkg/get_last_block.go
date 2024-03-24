package pkg

import (
	"fmt"
	"io"
	"net/http"
)

// GetLastBlock retrieves the latest block height from the mempool.space API.
func GetLastBlock() (string, error) {
	url := "https://mempool.space/api/blocks/tip/height"
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch last block: status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	lastBlock := string(body)
	return lastBlock, nil
}
