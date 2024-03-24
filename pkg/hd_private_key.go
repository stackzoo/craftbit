package pkg

import (
	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
)

// GenerateHDPrivateKey generates a new HD private key.
func GenerateHDPrivateKey() (string, error) {
	// Generate a random seed at the recommended length.
	seed, err := hdkeychain.GenerateSeed(hdkeychain.RecommendedSeedLen)
	if err != nil {
		return "", err
	}

	// Generate a new master node using the seed.
	key, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		return "", err
	}

	// Show that the generated master node extended key is private.
	base58PrivateKey := key.String()

	return base58PrivateKey, nil
}
