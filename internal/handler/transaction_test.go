package handler

import (
	"reflect"
	"tax-calculator/internal/models"
	"testing"
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
					Operation: "buy",
					UnitCost:  10.00,
					Quantity:  100,
				},
				{
					Code:      "A",
					Operation: "sell",
					UnitCost:  15.00,
					Quantity:  50,
				},
				{
					Code:      "A",
					Operation: "sell",
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
					Operation: "buy",
					UnitCost:  10.00,
					Quantity:  10000,
				},
				{
					Code:      "B",
					Operation: "sell",
					UnitCost:  20.00,
					Quantity:  5000,
				},
				{
					Code:      "B",
					Operation: "sell",
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
					Operation: "buy",
					UnitCost:  10.00,
					Quantity:  10000,
				},
				{
					Code:      "C",
					Operation: "buy",
					UnitCost:  25.00,
					Quantity:  5000,
				},
				{
					Code:      "C",
					Operation: "sell",
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
					Operation: "buy",
					UnitCost:  10.00,
					Quantity:  10000,
				},
				{
					Code:      "D",
					Operation: "buy",
					UnitCost:  25.00,
					Quantity:  5000,
				},
				{
					Code:      "D",
					Operation: "sell",
					UnitCost:  15.00,
					Quantity:  10000,
				},
				{
					Code:      "D",
					Operation: "sell",
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
					Operation: "buy",
					UnitCost:  10.00,
					Quantity:  100,
				},
				{
					Code:      "B",
					Operation: "buy",
					UnitCost:  10.00,
					Quantity:  10000,
				},

				{
					Code:      "A",
					Operation: "sell",
					UnitCost:  15.00,
					Quantity:  50,
				},
				{
					Code:      "B",
					Operation: "sell",
					UnitCost:  20.00,
					Quantity:  5000,
				},
				{
					Code:      "A",
					Operation: "sell",
					UnitCost:  15.00,
					Quantity:  50,
				},
				{
					Code:      "B",
					Operation: "sell",
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
			taxes, err := ProcessTransactions(test.transactions)
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
