package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"github.com/charmbracelet/lipgloss"
	"github.com/stackzoo/craftbit/pkg"
)

func main() {
	var utility string

	form := huh.NewForm(
		huh.NewGroup(huh.NewNote().
			Title("CraftBit").
			Description("Welcome to CraftBit!\n\nHow may we help you?")),

		// Choose a utility.
		huh.NewGroup(
			huh.NewSelect[string]().
				Options(huh.NewOptions("Decode Raw Transaction", "Bech32", "Hack Bitcoin")...).
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
			time.Sleep(2 * time.Second)
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

	case "Bech32":
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

	default:
		fmt.Println("Invalid selection")
	}

}

// #####################################################################################

// 	prepareBurger := func() {
// 		time.Sleep(2 * time.Second)
// 	}

// 	_ = spinner.New().Title("Preparing your burger...").Accessible(accessible).Action(prepareBurger).Run()

// 	// Print order summary.
// 	{
// 		var sb strings.Builder
// 		keyword := func(s string) string {
// 			return lipgloss.NewStyle().Foreground(lipgloss.Color("212")).Render(s)
// 		}
// 		fmt.Fprintf(&sb,
// 			"%s\n\nOne %s%s, topped with %s with %s on the side.",
// 			lipgloss.NewStyle().Bold(true).Render("BURGER RECEIPT"),
// 			keyword(order.Burger.Spice.String()),
// 			keyword(order.Burger.Type),
// 			keyword(xstrings.EnglishJoin(order.Burger.Toppings, true)),
// 			keyword(order.Side),
// 		)

// 		name := order.Name
// 		if name != "" {
// 			name = ", " + name
// 		}
// 		fmt.Fprintf(&sb, "\n\nThanks for your order%s!", name)

// 		if order.Discount {
// 			fmt.Fprint(&sb, "\n\nEnjoy 15% off.")
// 		}

// 		fmt.Println(
// 			lipgloss.NewStyle().
// 				Width(40).
// 				BorderStyle(lipgloss.RoundedBorder()).
// 				BorderForeground(lipgloss.Color("63")).
// 				Padding(1, 2).
// 				Render(sb.String()),
// 		)
// 	}
// }
