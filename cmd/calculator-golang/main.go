package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/ang-costa-neto/calculator-golang/internal/handler"
	"github.com/ang-costa-neto/calculator-golang/internal/models"
	"github.com/ang-costa-neto/calculator-golang/internal/parser"
)

func main() {
	var fileName string
	var inputJSON string
	var transactions []models.Transaction
	var err error

	flag.StringVar(&fileName, "file", "", "file containing JSON transactions")
	flag.StringVar(&inputJSON, "input", "", "JSON string containing transactions")
	flag.Parse()

	if fileName == "" && inputJSON == "" {
		fmt.Println("Please provide a file with transactions using -file flag or JSON input using -input flag.")
		os.Exit(1)
	}

	if inputJSON != "" {
		err = json.NewDecoder(strings.NewReader(inputJSON)).Decode(&transactions)
		if err != nil {
			fmt.Printf("Error decoding JSON input: %v\n", err)
			os.Exit(1)
		}
	} else if fileName != "" {
		transactions, err = parser.ParseTransactions(fileName)
		if err != nil {
			fmt.Printf("Error parsing transactions from file: %v\n", err)
			os.Exit(1)
		}
	}

	taxes, err := handler.ProcessTransactions(transactions)
	if err != nil {
		fmt.Printf("Error processing transactions: %v\n", err)
		os.Exit(1)
	}

	taxesJSON, err := json.Marshal(taxes)
	if err != nil {
		fmt.Printf("Error marshaling taxes to JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(taxesJSON))
}
