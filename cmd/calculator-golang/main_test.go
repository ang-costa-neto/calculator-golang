package main

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/ang-costa-neto/calculator-golang/internal/handler"
	"github.com/ang-costa-neto/calculator-golang/internal/models"
)

// Mock Processor
type mockProcessor struct {
	ProcessTransactionsFunc func(transactions []models.Transaction) ([]models.TaxResult, error)
}

func (m *mockProcessor) ProcessTransactions(transactions []models.Transaction) ([]models.TaxResult, error) {
	return m.ProcessTransactionsFunc(transactions)
}

func TestRunWithFileError(t *testing.T) {
	_, err := run("invalid_file_name.json", "")
	if err == nil {
		t.Fatalf("Expected error but got none")
	}
	if !strings.Contains(err.Error(), "error parsing transactions from file") {
		t.Errorf("Expected error message to contain 'error parsing transactions from file', but got %v", err)
	}
}

func TestRunWithInvalidJSON(t *testing.T) {
	invalidJSON := `[
		{"code": "A", "operation": "buy", "unit-cost": 10.00, "quantity": 100
	]`

	_, err := run("", invalidJSON)
	if err == nil {
		t.Fatalf("Expected error but got none")
	}
	if !strings.Contains(err.Error(), "error decoding JSON input") {
		t.Errorf("Expected error message to contain 'error decoding JSON input', but got %v", err)
	}
}

func TestRunWithProcessingError(t *testing.T) {
	mock := &mockProcessor{
		ProcessTransactionsFunc: func(transactions []models.Transaction) ([]models.TaxResult, error) {
			return nil, fmt.Errorf("mock processing error")
		},
	}
	handler.ProcessorInstance = mock

	jsonInput := `[
		{"code": "A", "operation": "buy", "unit-cost": 10.00, "quantity": 100}
	]`

	_, err := run("", jsonInput)
	if err == nil {
		t.Fatalf("Expected error but got none")
	}
	if !strings.Contains(err.Error(), "error processing transactions") {
		t.Errorf("Expected error message to contain 'error processing transactions', but got %v", err)
	}
}

func TestRunWithUnreadableFile(t *testing.T) {
	file, err := os.CreateTemp("", "transactions_*.json")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	file.Close()

	err = os.Chmod(file.Name(), 0000)
	if err != nil {
		t.Fatalf("Failed to change file permissions: %v", err)
	}
	defer func() {
		err = os.Chmod(file.Name(), 0644)
		if err != nil {
			t.Fatalf("Failed to restore file permissions: %v", err)
		}
		os.Remove(file.Name())
	}()

	_, err = run(file.Name(), "")
	if err == nil {
		t.Fatalf("Expected error but got none")
	}
	if !strings.Contains(err.Error(), "error parsing transactions from file") {
		t.Errorf("Expected error message to contain 'error parsing transactions from file', but got %v", err)
	}
}

// Helper function for comparing tax results
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
