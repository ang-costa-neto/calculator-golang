package parser

import (
	"reflect"
	"tax-calculator/internal/models"
	"testing"
)

func TestParseTransactions(t *testing.T) {
	tests := []struct {
		filename string
		expected []models.Transaction
	}{
		{
			filename: "../../transaction_file_mock/transaction1.json",
			expected: []models.Transaction{
				{
					Operation: "buy",
					UnitCost:  10.00,
					Quantity:  10000,
				},
				{
					Operation: "sell",
					UnitCost:  20.00,
					Quantity:  5000,
				},
				{
					Operation: "sell",
					UnitCost:  5.00,
					Quantity:  5000,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.filename, func(t *testing.T) {
			transaction, err := ParseTransactions(test.filename)

			if err != nil {
				t.Errorf("ParseTransactions() error = %v", err)
			}
			if !reflect.DeepEqual(transaction, test.expected) {
				t.Errorf("ParseTransactions() = %v, want %v", transaction, test.expected)
			}
		})
	}
}
