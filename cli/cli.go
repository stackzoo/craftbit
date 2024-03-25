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

func Run() {
	var utility string

	form := huh.NewForm(
		huh.NewGroup(huh.NewNote().
			Title("CraftBit").
			Description("Welcome to CraftBit!\n\nHow may we help you?")),

		// Choose a utility.
		huh.NewGroup(
			huh.NewSelect[string]().
				Options(huh.NewOptions("Decode Raw Transaction", "Bech32", "Transaction History", "Generate HD Private Key", "Generate P2PKH Address from HD Private Key",
					"Get Price", "Get Recommended Fees", "Get Last Block", "Get Lightning Network Stats", "Get Lightning Top Nodes", "Hack Bitcoin")...).
				Title("Choose your utility").
				Description("CraftBit has utilities for everyone!").
				Validate(func(t string) error {
					if t == "Hack Bitcoin" {
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

	switch utility {
	case "Decode Raw Transaction":
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
			fmt.Println("Uh oh:", err)
			os.Exit(1)
		}
		decodedTx, err := pkg.DecodeRawTransaction(rawTx)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		// Compose the output string
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
		decodeTx := func() {
			time.Sleep(1 * time.Second)
		}

		_ = spinner.New().Title("Decoding Transaction...").Action(decodeTx).Run()
		fmt.Println(
			lipgloss.NewStyle().
				Width(100).
				BorderStyle(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("63")).
				Padding(1, 2).
				Render(output.String()),
		)

	case "Transaction History":
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
			fmt.Println("Uh oh:", err)
			os.Exit(1)
		}

		transactionHistory, err := pkg.GetTransactionHistory(address)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		var transactionOutput string
		transactionOutput += fmt.Sprintf("Transaction History for %s\n", address)
		for i, tx := range transactionHistory {
			transactionOutput += fmt.Sprintf("%d. TxID: %s, Fee: %d\n", i+1, tx.TxID, tx.Amount)
		}

		TxHistory := func() {
			time.Sleep(1 * time.Second)
		}

		_ = spinner.New().Title("Retrieving latest transactions...").Action(TxHistory).Run()
		fmt.Println(
			lipgloss.NewStyle().
				Width(100).
				BorderStyle(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("63")).
				Padding(1, 2).
				Render(transactionOutput),
		)

	case "Bech32": // TO DO: COMPLETE LOGIC
		var operation string
		formBech32 := huh.NewForm(
			huh.NewGroup(huh.NewNote().
				Title("Bech32").
				Description("Bech32\n\nWhat Bech32 operaton do you want to perform?")),

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
			fmt.Println("Uh oh:", err)
			os.Exit(1)
		}

		switch operation {
		case "Encode String":
			hrp := "customHrp!11111q"
			data := []byte("Test data")
			encoded, err := pkg.EncodeBech32(hrp, data)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println("Encoded Data:", encoded)

			decodedHrp, decodedData, err := pkg.DecodeBech32(encoded)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println("Decoded HRP:", decodedHrp)
			fmt.Println("Decoded Data:", string(decodedData))
		}

	case "Generate HD Private Key":
		privateKey, err := pkg.GenerateHDPrivateKey()
		if err != nil {
			fmt.Println(err)
			return
		}
		generateSeed := func() {
			time.Sleep(1 * time.Second)
		}

		_ = spinner.New().Title("Generating Hierarchical Deterministic Private Key...").Action(generateSeed).Run()
		keyOutput := "BASE58 PRIVATE KEY:\n" + privateKey + "\n\nDO NOT SHARE WITH ANYONE 💀"
		fmt.Println(
			lipgloss.NewStyle().
				Width(100).
				BorderStyle(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("63")).
				Padding(1, 2).
				Render(keyOutput),
		)

	case "Generate P2PKH Address from HD Private Key":
		hdPrivateKey := ""
		formAddressFromHdPrivateKey := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Value(&hdPrivateKey).
					Title("Enter Bitcoin extended private key (xprv)").
					Placeholder("e.g., xprv9s21ZrQH143K2TQH119ScPgTzuL2mw2USC7QzsogsFobCeaSZ8Q8kUUMtkPWSksAEiDTSvRhtWkhk5AfZizxcWqx5NXDwkxKiJcnSGavNUZ").
					Description("Enter the Bitcoin HD Private Key for wich to generate a new address"),
			),
		)

		err := formAddressFromHdPrivateKey.Run()
		if err != nil {
			fmt.Println("Uh oh:", err)
			os.Exit(1)
		}

		key, err := pkg.CreateP2pkhAddressFromPrivateKey(hdPrivateKey)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		var transactionOutput string
		transactionOutput += fmt.Sprintf("Transaction History for %s\n", key)

		AddressFromPrivateKey := func() {
			time.Sleep(1 * time.Second)
		}

		_ = spinner.New().Title("Generating Address...").Action(AddressFromPrivateKey).Run()
		fmt.Println(
			lipgloss.NewStyle().
				Width(100).
				BorderStyle(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("63")).
				Padding(1, 2).
				Render("Generated P2PKH Address\n" + key),
		)

	case "Get Price":
		prices, err := pkg.GetPrices()
		if err != nil {
			fmt.Println(err)
			return
		}
		getPrices := func() {
			time.Sleep(1 * time.Second)
		}

		_ = spinner.New().Title("Get BTC current price...").Action(getPrices).Run()
		pricesDTO := StructFieldsToString(prices)
		fmt.Println(
			lipgloss.NewStyle().
				Width(100).
				BorderStyle(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("63")).
				Padding(1, 2).
				Render("PRICE\n\n" + pricesDTO),
		)

	case "Get Recommended Fees":
		fees, err := pkg.GetFees()
		if err != nil {
			fmt.Println(err)
			return
		}
		getFees := func() {
			time.Sleep(1 * time.Second)
		}

		_ = spinner.New().Title("Get recommended fees...").Action(getFees).Run()
		feesDTO := StructFieldsToString(fees)
		fmt.Println(
			lipgloss.NewStyle().
				Width(100).
				BorderStyle(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("63")).
				Padding(1, 2).
				Render("RECOMMENDED FEES\n\n" + feesDTO),
		)

	case "Get Lightning Network Stats":
		lnStats, err := pkg.GetLightningStatistics()
		if err != nil {
			fmt.Println(err)
			return
		}
		getLnStats := func() {
			time.Sleep(1 * time.Second)
		}

		_ = spinner.New().Title("Get lightning network statistics...").Action(getLnStats).Run()
		lnStatsDTO := StructFieldsToString(lnStats.Latest)
		fmt.Println(
			lipgloss.NewStyle().
				Width(100).
				BorderStyle(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("63")).
				Padding(1, 2).
				Render("LIGHTNING NETWORK STATS\n\n" + lnStatsDTO),
		)

	case "Get Lightning Top Nodes":
		lnTopNodes, err := pkg.GetLightningTopNodes()
		if err != nil {
			fmt.Println(err)
			return
		}
		getLnTopNodes := func() {
			time.Sleep(1 * time.Second)
		}

		_ = spinner.New().Title("Get lightning network top nodes...").Action(getLnTopNodes).Run()
		lnTopNodesCapacity := lnTopNodes.ByCapacityString()
		fmt.Println(
			lipgloss.NewStyle().
				Width(100).
				BorderStyle(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("63")).
				Padding(1, 2).
				Render("LIGHTNING TOP NODES BY CAPACITY\n\n" + lnTopNodesCapacity),
		)
		lnTopNodesChannels := lnTopNodes.ByChannelsString()
		fmt.Println(
			lipgloss.NewStyle().
				Width(100).
				BorderStyle(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("63")).
				Padding(1, 2).
				Render("LIGHTNING TOP NODES BY NUMBER OF CHANNELS\n\n" + lnTopNodesChannels),
		)

	case "Get Last Block":
		lastBlock, err := pkg.GetLastBlock()
		if err != nil {
			fmt.Println(err)
			return
		}
		getLastBlock := func() {
			time.Sleep(1 * time.Second)
		}

		_ = spinner.New().Title("Get last block in the blockchain...").Action(getLastBlock).Run()
		fmt.Println(
			lipgloss.NewStyle().
				Width(100).
				BorderStyle(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("63")).
				Padding(1, 2).
				Render("LAST BLOCK\n\n" + lastBlock),
		)

	default:
		fmt.Println("Invalid selection")
	}

}

func PrintStructFields(s interface{}) {
	val := reflect.ValueOf(s)
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := typ.Field(i).Name
		fieldValue := field.Interface()

		fmt.Printf("%s: %v\n", fieldName, fieldValue)
	}
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
