package main

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"

	"github.com/ang-costa-neto/calculator-golang/internal/models"
)

func TestMainWithFile(t *testing.T) {
	// JSON de transações para teste
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
	// Criar arquivo temporário
	file, err := os.CreateTemp("", "transactions_*.json")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(file.Name())

	_, err = file.WriteString(jsonData)
	if err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}
	file.Close()

	// Capturar saída padrão
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Simular argumento de linha de comando com -file
	os.Args = []string{"main.go", "-file", file.Name()}
	main()

	// Restaurar saída padrão e ler saída capturada
	w.Close()
	os.Stdout = old

	var capturedOutput bytes.Buffer
	_, err = capturedOutput.ReadFrom(r)
	if err != nil {
		t.Fatalf("Failed to read captured output: %v", err)
	}

	// Desserializar saída para comparação
	var actualOutput []models.TaxResult
	err = json.Unmarshal(capturedOutput.Bytes(), &actualOutput)
	if err != nil {
		t.Fatalf("Failed to unmarshal captured output: %v", err)
	}

	// Saída esperada
	expectedOutput := []models.TaxResult{
		{Code: "A", Tax: 0.00},
		{Code: "A", Tax: 0.00},
		{Code: "A", Tax: 0.00},
	}

	// Comparação dos resultados
	if !equalTaxResults(actualOutput, expectedOutput) {
		t.Errorf("Expected output %v, but got %v", expectedOutput, actualOutput)
	}
}

func TestMainWithJSONInput(t *testing.T) {
	// JSON de transações para teste
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

	// Capturar saída padrão
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Simular argumento de linha de comando com -input
	os.Args = []string{"main.go", "-input", jsonInput}
	main()

	// Restaurar saída padrão e ler saída capturada
	w.Close()
	os.Stdout = old

	var capturedOutput bytes.Buffer
	_, err := capturedOutput.ReadFrom(r)
	if err != nil {
		t.Fatalf("Failed to read captured output: %v", err)
	}

	// Desserializar saída para comparação
	var actualOutput []models.TaxResult
	err = json.Unmarshal(capturedOutput.Bytes(), &actualOutput)
	if err != nil {
		t.Fatalf("Failed to unmarshal captured output: %v", err)
	}

	// Saída esperada
	expectedOutput := []models.TaxResult{
		{Code: "A", Tax: 0.00},
		{Code: "A", Tax: 0.00},
		{Code: "A", Tax: 0.00},
	}

	// Comparação dos resultados
	if !equalTaxResults(actualOutput, expectedOutput) {
		t.Errorf("Expected output %v, but got %v", expectedOutput, actualOutput)
	}
}

// Função auxiliar para comparar os resultados do imposto
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
