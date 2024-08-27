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

func run(fileName string, inputJSON string) ([]models.TaxResult, error) {
	if (fileName == "" && inputJSON == "") || (fileName != "" && inputJSON != "") {
		return nil, fmt.Errorf("please provide either a file with transactions using -file flag or a JSON input using -input flag, but not both")
	}

	var transactions []models.Transaction
	var err error

	if inputJSON != "" {
		err = json.NewDecoder(strings.NewReader(inputJSON)).Decode(&transactions)
		if err != nil {
			return nil, fmt.Errorf("error decoding JSON input: %w", err)
		}
	} else if fileName != "" {
		transactions, err = parser.ParseTransactions(fileName)
		if err != nil {
			return nil, fmt.Errorf("error parsing transactions from file: %w", err)
		}
	}

	taxes, err := handler.ProcessorInstance.ProcessTransactions(transactions)
	if err != nil {
		return nil, fmt.Errorf("error processing transactions: %w", err)
	}

	return taxes, nil
}

func main() {
	var fileName string
	var inputJSON string

	flag.StringVar(&fileName, "file", "", "file containing JSON transactions")
	flag.StringVar(&inputJSON, "input", "", "JSON string containing transactions")
	flag.Parse()

	taxes, err := run(fileName, inputJSON)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	taxesJSON, err := json.Marshal(taxes)
	if err != nil {
		fmt.Printf("Error marshaling taxes to JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(taxesJSON))
}
