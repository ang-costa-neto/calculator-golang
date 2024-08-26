package handler

import (
	"reflect"
	"testing"

	"github.com/ang-costa-neto/calculator-golang/internal/constants"
	"github.com/ang-costa-neto/calculator-golang/internal/models"
)

func TestProcessTransactions(t *testing.T) {
	tests := []struct {
		name         string
		transactions []models.Transaction
		expected     []models.TaxResult
	}{
		{
			name: "Case A",
			transactions: []models.Transaction{
				{
					Code:      "A",
					Operation: constants.BuyOperation,
					UnitCost:  10.00,
					Quantity:  100,
				},
				{
					Code:      "A",
					Operation: constants.SellOperation,
					UnitCost:  15.00,
					Quantity:  50,
				},
				{
					Code:      "A",
					Operation: constants.SellOperation,
					UnitCost:  15.00,
					Quantity:  50,
				},
			},
			expected: []models.TaxResult{
				{
					Code: "A",
					Tax:  0.00,
				},
				{
					Code: "A",
					Tax:  0.00,
				},
				{
					Code: "A",
					Tax:  0.00,
				},
			},
		},
		{
			name: "Case B",
			transactions: []models.Transaction{
				{
					Code:      "B",
					Operation: constants.BuyOperation,
					UnitCost:  10.00,
					Quantity:  10000,
				},
				{
					Code:      "B",
					Operation: constants.SellOperation,
					UnitCost:  20.00,
					Quantity:  5000,
				},
				{
					Code:      "B",
					Operation: constants.SellOperation,
					UnitCost:  5.00,
					Quantity:  5000,
				},
			},
			expected: []models.TaxResult{
				{
					Code: "B",
					Tax:  0.00,
				},
				{
					Code: "B",
					Tax:  10000.00,
				},
				{
					Code: "B",
					Tax:  0.00,
				},
			},
		},
		{
			name: "Case C",
			transactions: []models.Transaction{
				{
					Code:      "C",
					Operation: constants.BuyOperation,
					UnitCost:  10.00,
					Quantity:  10000,
				},
				{
					Code:      "C",
					Operation: constants.BuyOperation,
					UnitCost:  25.00,
					Quantity:  5000,
				},
				{
					Code:      "C",
					Operation: constants.SellOperation,
					UnitCost:  15.00,
					Quantity:  10000,
				},
			},
			expected: []models.TaxResult{
				{
					Code: "C",
					Tax:  0.00,
				},
				{
					Code: "C",
					Tax:  0.00,
				},
				{
					Code: "C",
					Tax:  0.00,
				},
			},
		},
		{
			name: "Case D",
			transactions: []models.Transaction{
				{
					Code:      "D",
					Operation: constants.BuyOperation,
					UnitCost:  10.00,
					Quantity:  10000,
				},
				{
					Code:      "D",
					Operation: constants.BuyOperation,
					UnitCost:  25.00,
					Quantity:  5000,
				},
				{
					Code:      "D",
					Operation: constants.SellOperation,
					UnitCost:  15.00,
					Quantity:  10000,
				},
				{
					Code:      "D",
					Operation: constants.SellOperation,
					UnitCost:  25.00,
					Quantity:  5000,
				},
			},
			expected: []models.TaxResult{
				{
					Code: "D",
					Tax:  0.00,
				},
				{
					Code: "D",
					Tax:  0.00,
				},
				{
					Code: "D",
					Tax:  0.00,
				},
				{
					Code: "D",
					Tax:  10000.00,
				},
			},
		},
		{
			name: "Mixed Case",
			transactions: []models.Transaction{
				{
					Code:      "A",
					Operation: constants.BuyOperation,
					UnitCost:  10.00,
					Quantity:  100,
				},
				{
					Code:      "B",
					Operation: constants.BuyOperation,
					UnitCost:  10.00,
					Quantity:  10000,
				},
				{
					Code:      "A",
					Operation: constants.SellOperation,
					UnitCost:  15.00,
					Quantity:  50,
				},
				{
					Code:      "B",
					Operation: constants.SellOperation,
					UnitCost:  20.00,
					Quantity:  5000,
				},
				{
					Code:      "A",
					Operation: constants.SellOperation,
					UnitCost:  15.00,
					Quantity:  50,
				},
				{
					Code:      "B",
					Operation: constants.SellOperation,
					UnitCost:  5.00,
					Quantity:  5000,
				},
			},
			expected: []models.TaxResult{
				{
					Code: "A",
					Tax:  0.00,
				},
				{
					Code: "B",
					Tax:  0.00,
				},
				{
					Code: "A",
					Tax:  0.00,
				},
				{
					Code: "B",
					Tax:  10000.00,
				},
				{
					Code: "A",
					Tax:  0.00,
				},
				{
					Code: "B",
					Tax:  0.00,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Use a real processor instance
			processor := ProcessorInstance
			taxes, err := processor.ProcessTransactions(test.transactions)
			if err != nil {
				t.Error(err)
				return
			}
			if !reflect.DeepEqual(test.expected, taxes) {
				t.Errorf("expected: %v, got: %v", test.expected, taxes)
			}
		})
	}
}
