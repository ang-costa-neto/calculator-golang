package parser

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/ang-costa-neto/calculator-golang/internal/models"
)

func TestParseTransactions(t *testing.T) {
	tests := []struct {
		name     string
		fileName string
		expected []models.Transaction
		wantErr  bool
	}{
		{
			name:     "Valid File",
			fileName: "../../transaction_file_mock/transaction1.json",
			expected: []models.Transaction{
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
			wantErr: false,
		},
		{
			name:     "File Not Found",
			fileName: "../../transaction_file_mock/nonexistent.json",
			expected: nil,
			wantErr:  true,
		},
		{
			name:     "Invalid JSON",
			fileName: "../../transaction_file_mock/invalid.json",
			expected: nil,
			wantErr:  true,
		},
	}

	// Create temporary files for testing invalid JSON
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "Invalid JSON" {
				// Create a temporary file with invalid JSON
				file, err := ioutil.TempFile("../../transaction_file_mock", "invalid-*.json")
				if err != nil {
					t.Fatalf("Failed to create temp file: %v", err)
				}
				defer os.Remove(file.Name())

				if _, err := file.WriteString("{invalid-json}"); err != nil {
					t.Fatalf("Failed to write to temp file: %v", err)
				}

				file.Close()
				tt.fileName = file.Name()
			}

			if tt.wantErr {
				_, err := ParseTransactions(tt.fileName)
				if err == nil {
					t.Errorf("Expected an error but got none")
				}
				return
			}

			transaction, err := ParseTransactions(tt.fileName)
			if err != nil {
				t.Errorf("ParseTransactions() error = %v", err)
				return
			}
			if !reflect.DeepEqual(transaction, tt.expected) {
				t.Errorf("ParseTransactions() = %v, want %v", transaction, tt.expected)
			}
		})
	}
}
