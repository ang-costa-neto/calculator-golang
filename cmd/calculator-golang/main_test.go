package main

import (
	"os"
	"testing"

	"github.com/ang-costa-neto/calculator-golang/internal/models"
)

func TestMainWithFile(t *testing.T) {
	jsonData := `
		[
			{
				"code":      "A",
				"operation": "buy",
				"unit-cost":  10.00,
				"quantity":  100
			},
			{
				"code":      "A",
				"operation": "sell",
				"unit-cost":  15.00,
				"quantity":  50
			},
			{
				"code":      "A",
				"operation": "sell",
				"unit-cost":  15.00,
				"quantity":  50
			}
		]
	`

	file, err := os.CreateTemp("", "transactions_*.json")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer func() {
		file.Close()
		os.Remove(file.Name())
	}()

	_, err = file.WriteString(jsonData)
	if err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}

	taxes, err := run(file.Name(), "")
	if err != nil {
		t.Fatalf("Error running with file: %v", err)
	}

	expectedOutput := []models.TaxResult{
		{Code: "A", Tax: 0.00},
		{Code: "A", Tax: 0.00},
		{Code: "A", Tax: 0.00},
	}

	if !equalTaxResults(taxes, expectedOutput) {
		t.Errorf("Expected output %v, but got %v", expectedOutput, taxes)
	}
}

func TestMainWithJSONInput(t *testing.T) {
	jsonInput := `
		[
			{
				"code":      "A",
				"operation": "buy",
				"unit-cost":  10.00,
				"quantity":  100
			},
			{
				"code":      "A",
				"operation": "sell",
				"unit-cost":  15.00,
				"quantity":  50
			},
			{
				"code":      "A",
				"operation": "sell",
				"unit-cost":  15.00,
				"quantity":  50
			}
		]
	`

	taxes, err := run("", jsonInput)
	if err != nil {
		t.Fatalf("Error running with JSON input: %v", err)
	}

	expectedOutput := []models.TaxResult{
		{Code: "A", Tax: 0.00},
		{Code: "A", Tax: 0.00},
		{Code: "A", Tax: 0.00},
	}

	if !equalTaxResults(taxes, expectedOutput) {
		t.Errorf("Expected output %v, but got %v", expectedOutput, taxes)
	}
}

func equalTaxResults(a, b []models.TaxResult) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].Code != b[i].Code || a[i].Tax != b[i].Tax {
			return false
		}
	}
	return true
}
