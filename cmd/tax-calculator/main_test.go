package tax_calculator

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestMainWithFile(t *testing.T) {
	jsonData := `[{"operation":"buy","unit-cost":10.00,"quantity":100},{"operation":"sell","unit-cost":15.00,"quantity":50},{"operation":"sell","unit-cost":15.00,"quantity":50}]`
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

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	os.Args = []string{"main.go", "-file", file.Name()}
	main()

	w.Close()
	os.Stdout = old

	var capturedOutput bytes.Buffer
	_, err = capturedOutput.ReadFrom(r)
	if err != nil {
		t.Fatalf("Failed to read captured output: %v", err)
	}

	expectedOutput := `[{"tax":0},{"tax":0},{"tax":0}]`
	if strings.TrimSpace(capturedOutput.String()) != expectedOutput {
		t.Errorf("Expected output %q, but got %q", expectedOutput, capturedOutput.String())
	}
}

func TestMainWithJSONInput(t *testing.T) {
	jsonInput := `[{"operation":"buy","unit-cost":10.00,"quantity":100},{"operation":"sell","unit-cost":15.00,"quantity":50},{"operation":"sell","unit-cost":15.00,"quantity":50}]`

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	os.Args = []string{"main.go", "-input", jsonInput}
	main()

	w.Close()
	os.Stdout = old

	var capturedOutput bytes.Buffer
	_, err := capturedOutput.ReadFrom(r)
	if err != nil {
		t.Fatalf("Failed to read captured output: %v", err)
	}

	expectedOutput := `[{"tax":0},{"tax":0},{"tax":0}]`
	if strings.TrimSpace(capturedOutput.String()) != expectedOutput {
		t.Errorf("Expected output %q, but got %q", expectedOutput, capturedOutput.String())
	}
}
