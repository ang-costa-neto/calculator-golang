package handler

import (
	"reflect"
	"tax-calculator/internal/models"
	"testing"
)

func TestProcessTransactions(t *testing.T) {
	tests := []struct {
		transactions []models.Transaction
		expected     []models.TaxResult
	}{
		{
			transactions: []models.Transaction{
				{
					Operation: "buy",
					UnitCost:  10.00,
					Quantity:  100,
				},
				{
					Operation: "sell",
					UnitCost:  15.00,
					Quantity:  50,
				},
				{
					Operation: "sell",
					UnitCost:  15.00,
					Quantity:  50,
				},
			},
			expected: []models.TaxResult{
				{
					Tax: 0.00,
				},
				{
					Tax: 0.00,
				},
				{
					Tax: 0.00,
				},
			},
		},
		{
			transactions: []models.Transaction{
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
			expected: []models.TaxResult{
				{
					Tax: 0.00,
				},
				{
					Tax: 10000.00,
				},
				{
					Tax: 0.00,
				},
			},
		},
		{
			transactions: []models.Transaction{
				{
					Operation: "buy",
					UnitCost:  10.00,
					Quantity:  10000,
				},
				{
					Operation: "buy",
					UnitCost:  25.00,
					Quantity:  5000,
				},
				{
					Operation: "sell",
					UnitCost:  15.00,
					Quantity:  10000,
				},
			},
			expected: []models.TaxResult{
				{
					Tax: 0.00,
				},
				{
					Tax: 0.00,
				},
				{
					Tax: 0.00,
				},
			},
		},
		{
			transactions: []models.Transaction{
				{
					Operation: "buy",
					UnitCost:  10.00,
					Quantity:  10000,
				},
				{
					Operation: "buy",
					UnitCost:  25.00,
					Quantity:  5000,
				},
				{
					Operation: "sell",
					UnitCost:  15.00,
					Quantity:  10000,
				},
				{
					Operation: "sell",
					UnitCost:  25.00,
					Quantity:  5000,
				},
			},
			expected: []models.TaxResult{
				{
					Tax: 0.00,
				},
				{
					Tax: 0.00,
				},
				{
					Tax: 0.00,
				},
				{
					Tax: 10000.00,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
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
