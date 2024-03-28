package cli

import (
	"encoding/hex"
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"github.com/charmbracelet/lipgloss"
	"github.com/stackzoo/craftbit/pkg"
)

type UtilityRunner struct {
	Utility   string
	Operation func() (interface{}, error)
	Spinner   string
}

var utilities = []UtilityRunner{
	{"Decode Raw Transaction", decodeRawTransaction, "Decoding Transaction..."},
	{"Transaction History", transactionHistory, "Retrieving latest transactions..."},
	{"Bech32", bech32Operation, ""},
	{"Generate HD Private Key", generateHDPrivateKey, "Generating Hierarchical Deterministic Private Key..."},
	{"Generate P2PKH Address from HD Private Key", generateP2PKHAddress, "Generating Address..."},
	{"Get Price", getPrice, "Get BTC current price..."},
	{"Get Recommended Fees", getRecommendedFees, "Get recommended fees..."},
	{"Get Lightning Network Stats", getLightningNetworkStats, "Get lightning network statistics..."},
	{"Get Lightning Top Nodes", getLightningTopNodes, "Get lightning network top nodes..."},
	{"Get Last Block", getLastBlock, "Get last block in the blockchain..."},
	{"Get Mempool Status", getMempoolStatus, "Get mempool status..."},
	{"Get Current Hashrate", getCurrentHashrate, "Get current hashrate..."},
	{"Hack Bitcoin 👾", nil, ""},
}

func Run() {
	for {
		utility := chooseUtility()
		if utility == "" {
			break
		}

		for _, u := range utilities {
			if u.Utility == utility {
				runUtility(u)
				break
			}
		}

		if !returnToMainForm() {
			break
		}
	}
}

func chooseUtility() string {
	var utility string

	form := huh.NewForm(
		huh.NewGroup(huh.NewNote().
			Title("CraftBit").
			Description("Welcome to CraftBit, the Bitcoin swiss army knife 🪛\n\nHow may we help you?")),

		// Choose a utility.
		huh.NewGroup(
			huh.NewSelect[string]().
				Options(huh.NewOptions(getUtilityNames()...)...).
				Title("Choose your utility").
				Description("CraftBit has utilities for everyone!").
				Validate(func(t string) error {
					if t == "Hack Bitcoin 👾" {
						return fmt.Errorf("nice try, cannot do that, sorry 😊")
					}
					return nil
				}).
				Value(&utility),
		),
	)

	err := form.Run()
	if err != nil {
		fmt.Println("Uh oh:", err)
		os.Exit(1)
	}

	return utility
}

func getUtilityNames() []string {
	names := make([]string, len(utilities))
	for i, u := range utilities {
		names[i] = u.Utility
	}
	return names
}

func runUtility(utility UtilityRunner) {
	result, err := utility.Operation()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if utility.Spinner != "" {
		runSpinner(utility.Spinner)
	}

	printResult(result)
}

func runSpinner(msg string) {
	decodeTx := func() {
		time.Sleep(1 * time.Second)
	}

	_ = spinner.New().Title(msg).Action(decodeTx).Run()
}

func printResult(result interface{}) {
	// Perform type assertion to convert result to string
	strResult, ok := result.(string)
	if !ok {
		fmt.Println("Error: result is not a string")
		return
	}

	fmt.Println(
		lipgloss.NewStyle().
			Width(100).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("63")).
			Padding(1, 2).
			Render(strResult),
	)
}

func returnToMainForm() bool {
	var returnToMainForm bool
	formReturn := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("Return to main form?").
				Value(&returnToMainForm).
				Description("Do you want to select another Bitcoin utility?"),
		),
	)

	err := formReturn.Run()
	if err != nil {
		fmt.Println("Uh oh:", err)
		os.Exit(1)
	}

	return returnToMainForm
}

func decodeRawTransaction() (interface{}, error) {
	rawTx := ""

	formDecodeRawTransaction := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Value(&rawTx).
				Title("Paste the full raw transaction here").
				Placeholder(".....").
				Description("This is gonna be a loooong string"),
		),
	)

	err := formDecodeRawTransaction.Run()
	if err != nil {
		return nil, err
	}

	decodedTx, err := pkg.DecodeRawTransaction(rawTx)
	if err != nil {
		return nil, err
	}

	var output strings.Builder
	output.WriteString(fmt.Sprintf("TRANSACTION ID: %s\n", decodedTx.TxID))
	for i, input := range decodedTx.Inputs {
		inputStr := fmt.Sprintf("Previous Tx: %s, Output Index: %d, ScriptSig: %s\n", input.PreviousOutPoint.Hash.String(), input.PreviousOutPoint.Index, hex.EncodeToString(input.SignatureScript))
		output.WriteString(fmt.Sprintf("TRANSACTION INPUT %d: %s\n", i+1, inputStr))
	}
	for i, out := range decodedTx.Outputs {
		outputStr := fmt.Sprintf("Value: %d, ScriptPubKey: %s\n", out.Value, hex.EncodeToString(out.PkScript))
		output.WriteString(fmt.Sprintf("TRANSACTION OUTPUT %d: %s\n", i+1, outputStr))
	}

	return output.String(), nil
}

func transactionHistory() (interface{}, error) {
	address := ""

	formTransactionHistory := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Value(&address).
				Title("Enter Bitcoin address").
				Placeholder("e.g., bc1q0xs2775td0fm2t80m5alnmv5j6jhxqkgsdz5rv").
				Description("Enter the Bitcoin address to fetch transaction history"),
		),
	)

	err := formTransactionHistory.Run()
	if err != nil {
		return nil, err
	}

	transactionHistory, err := pkg.GetTransactionHistory(address)
	if err != nil {
		return nil, err
	}

	var transactionOutput strings.Builder
	transactionOutput.WriteString(fmt.Sprintf("Transaction History for %s\n", address))
	for i, tx := range transactionHistory {
		transactionOutput.WriteString(fmt.Sprintf("%d. TxID: %s, Fee: %d\n", i+1, tx.TxID, tx.Amount))
	}

	return transactionOutput.String(), nil
}

func bech32Operation() (interface{}, error) {
	var operation string
	formBech32 := huh.NewForm(
		huh.NewGroup(huh.NewNote().
			Title("Bech32").
			Description("Bech32\n\nWhat Bech32 operation do you want to perform?")),

		// Choose a Bech32 operation.
		huh.NewGroup(
			huh.NewSelect[string]().
				Options(huh.NewOptions("Decode String", "Encode String")...).
				Title("Choose your operation").
				Value(&operation),
		),
	)

	err := formBech32.Run()
	if err != nil {
		return nil, err
	}

	switch operation {
	case "Encode String":
		hrp := "customHrp!11111q"
		data := []byte("Test data")
		encoded, err := pkg.EncodeBech32(hrp, data)
		if err != nil {
			return nil, err
		}
		fmt.Println("Encoded Data:", encoded)

		decodedHrp, decodedData, err := pkg.DecodeBech32(encoded)
		if err != nil {
			return nil, err
		}
		fmt.Println("Decoded HRP:", decodedHrp)
		fmt.Println("Decoded Data:", string(decodedData))
	}

	return nil, nil
}

func generateHDPrivateKey() (interface{}, error) {
	privateKey, err := pkg.GenerateHDPrivateKey()
	if err != nil {
		return nil, err
	}

	keyOutput := "BASE58 PRIVATE KEY:\n" + privateKey + "\n\nDO NOT SHARE WITH ANYONE 💀"
	return keyOutput, nil
}

func generateP2PKHAddress() (interface{}, error) {
	hdPrivateKey := ""
	formAddressFromHdPrivateKey := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Value(&hdPrivateKey).
				Title("Enter Bitcoin extended private key (xprv)").
				Placeholder("e.g., xprv9s21ZrQH143K2TQH119ScPgTzuL2mw2USC7QzsogsFobCeaSZ8Q8kUUMtkPWSksAEiDTSvRhtWkhk5AfZizxcWqx5NXDwkxKiJcnSGavNUZ").
				Description("Enter the Bitcoin HD Private Key for which to generate a new address"),
		),
	)

	err := formAddressFromHdPrivateKey.Run()
	if err != nil {
		return nil, err
	}

	key, err := pkg.CreateP2pkhAddressFromPrivateKey(hdPrivateKey)
	if err != nil {
		return nil, err
	}

	return "Generated P2PKH Address\n" + key, nil
}

func getPrice() (interface{}, error) {
	prices, err := pkg.GetPrices()
	if err != nil {
		return nil, err
	}

	return StructFieldsToString(prices), nil
}

func getRecommendedFees() (interface{}, error) {
	fees, err := pkg.GetFees()
	if err != nil {
		return nil, err
	}

	return StructFieldsToString(fees), nil
}

func getLightningNetworkStats() (interface{}, error) {
	lnStats, err := pkg.GetLightningStatistics()
	if err != nil {
		return nil, err
	}

	return StructFieldsToString(lnStats.Latest), nil
}

func getLightningTopNodes() (interface{}, error) {
	lnTopNodes, err := pkg.GetLightningTopNodes()
	if err != nil {
		return nil, err
	}

	lnTopNodesCapacity := lnTopNodes.ByCapacityString()
	lnTopNodesChannels := lnTopNodes.ByChannelsString()

	return "LIGHTNING TOP NODES BY CAPACITY\n\n" + lnTopNodesCapacity + "\n\nLIGHTNING TOP NODES BY NUMBER OF CHANNELS\n\n" + lnTopNodesChannels, nil
}

func getLastBlock() (interface{}, error) {
	lastBlock, err := pkg.GetLastBlock()
	if err != nil {
		return nil, err
	}

	return "LAST BLOCK\n\n" + lastBlock, nil
}

func getMempoolStatus() (interface{}, error) {
	mempoolStatus, err := pkg.GetMempoolStatus()
	if err != nil {
		return nil, err
	}

	return StructFieldsToString(mempoolStatus), nil
}

func getCurrentHashrate() (interface{}, error) {
	currentHashrate, err := pkg.GetCurrentHashrate()
	if err != nil {
		return nil, err
	}

	return StructFieldsToString(currentHashrate), nil
}

func StructFieldsToString(s interface{}) string {
	val := reflect.ValueOf(s)
	typ := val.Type()

	var builder strings.Builder

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := typ.Field(i).Name
		fieldValue := fmt.Sprintf("%v", field.Interface())

		builder.WriteString(fmt.Sprintf("%s: %s\n", fieldName, fieldValue))
	}

	return builder.String()
}
