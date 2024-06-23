package parser

import (
	"encoding/json"
	"os"
	"tax-calculator/internal/models"
)

func ParseTransactions(fileName string) ([]models.Transaction, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	var transactions []models.Transaction
	err = json.NewDecoder(file).Decode(&transactions)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}
