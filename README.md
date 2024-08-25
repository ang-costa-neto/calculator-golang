![coverage](https://img.shields.io/codecov/c/github/ang-costa-neto/calculator-golang)

# Tax Calculator

This project is a simple tax calculator implemented in Go.

## How to Build

To build the Tax Calculator, run command:

```bash
   make build
```

## How to run

Once you have built the project, you can run the Tax Calculator:

### Linux / macOS

Navigate to the directory containing the executable and run:

```bash
./calculator-golang
```

### Windows

Navigate to the directory containing the executable and run:

```bash
calculator-golang.exe
```

## Usage

The input file must be in JSON format as shown below:

```json
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
```
### Running with a file
To run the Tax Calculator using an input file, use the following command:
```bash
./calculator-golang -file transaction.json
```
Replace transaction.json with the actual input file containing your transaction data.

### Running with JSON Input Directly
To pass JSON data directly via the command line, use the following command:
```bash
./calculator-golang -input '[{"code":"A","operation": "buy","unit-cost":10.00,"quantity":100},{"code":"A","operation": "sell","unit-cost":15.00,"quantity":50},{"code":"A","operation": "sell","unit-cost":15.00,"quantity":50}]'
```
Make sure the JSON string is correctly formatted and enclosed in single quotes to avoid issues with shell parsing.

