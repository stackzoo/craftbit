package pkg

import (
	"fmt"

	"github.com/btcsuite/btcd/btcutil/bech32"
)

// EncodeBech32 encodes the given data using the Bech32 encoding with the specified human-readable part.
func EncodeBech32(hrp string, data []byte) (string, error) {
	// Convert data to base32
	conv, err := bech32.ConvertBits(data, 8, 5, true)
	if err != nil {
		return "", fmt.Errorf("failed to convert bits: %v", err)
	}

	// Encode using Bech32
	encoded, err := bech32.Encode(hrp, conv)
	if err != nil {
		return "", fmt.Errorf("failed to encode using Bech32: %v", err)
	}

	return encoded, nil
}

// DecodeBech32 decodes the given Bech32 encoded string and returns the decoded data and human-readable part.
func DecodeBech32(encoded string) (string, []byte, error) {
	hrp, data, err := bech32.Decode(encoded)
	if err != nil {
		return "", nil, fmt.Errorf("failed to decode Bech32 string: %v", err)
	}

	// Convert data back from base32
	conv, err := bech32.ConvertBits(data, 5, 8, false)
	if err != nil {
		return "", nil, fmt.Errorf("failed to convert bits: %v", err)
	}

	return hrp, conv, nil
}
