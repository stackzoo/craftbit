package pkg

import (
	"bytes"
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/wire"
)

// DecodedTransaction represents the decoded Bitcoin transaction.
type DecodedTransaction struct {
	TxID    string
	Inputs  []*wire.TxIn
	Outputs []*wire.TxOut
}

// DecodeRawTransaction decodes a raw Bitcoin transaction and returns the decoded transaction information.
func DecodeRawTransaction(rawTx string) (*DecodedTransaction, error) {
	// Decode the hexadecimal string into bytes
	rawBytes, err := hex.DecodeString(rawTx)
	if err != nil {
		return nil, fmt.Errorf("failed to decode raw transaction: %v", err)
	}

	// Deserialize the bytes into a MsgTx struct
	var msgTx wire.MsgTx
	err = msgTx.Deserialize(bytes.NewReader(rawBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to deserialize raw transaction: %v", err)
	}

	// Get the transaction ID
	txID := msgTx.TxHash().String()

	return &DecodedTransaction{
		TxID:    txID,
		Inputs:  msgTx.TxIn,
		Outputs: msgTx.TxOut,
	}, nil
}
