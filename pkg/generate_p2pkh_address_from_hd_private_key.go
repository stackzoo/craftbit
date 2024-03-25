package pkg

import (
	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
)

// Create public address from an HD private key.
func CreateP2pkhAddressFromPrivateKey(base58PrivateKey string) (string, error) {

	// Parse the decoded key.
	key, err := hdkeychain.NewKeyFromString(base58PrivateKey)
	if err != nil {
		return "", err
	}
	address, err := key.Address(&chaincfg.MainNetParams)
	if err != nil {
		return "", err
	}
	return address.String(), nil
}
