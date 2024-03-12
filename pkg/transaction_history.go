// pkg/transaction_history.go

package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Transaction represents information about a single transaction.
type Transaction struct {
	TxID   string `json:"txid"`
	Amount int64  `json:"fee"`
}

// GetTransactionHistory retrieves the transaction history of a Bitcoin address from mempool.space API.
func GetTransactionHistory(address string) ([]Transaction, error) {
	url := fmt.Sprintf("https://mempool.space/api/address/%s/txs", address)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch transaction history: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch transaction history: status code %d", resp.StatusCode)
	}

	var transactions []Transaction
	err = json.NewDecoder(resp.Body).Decode(&transactions)
	if err != nil {
		return nil, fmt.Errorf("failed to decode transaction history response: %v", err)
	}

	return transactions, nil
}
