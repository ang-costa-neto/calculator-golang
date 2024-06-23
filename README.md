# Tax Calculator

This project is a simple tax calculator implemented in Go.

## How to Build

To build the Tax Calculator, follow these steps:

1. Navigate to the `cmd/tax-calculator` directory:
   ```bash
    cd cmd/tax-calculator
   ```
2. Build the main executable:
    ```
    go build
    ```
   This will create an executable file named tax-calculator (or tax-calculator.exe on Windows) in the current directory.

## How to run

Once you have built the project, you can run the Tax Calculator:

### Linux / macOS

Navigate to the directory containing the executable and run:

```bash
./tax-calculator
```

### Windows

Navigate to the directory containing the executable and run:

```bash
tax-calculator.exe
```

## Usage

The input file must be in JSON format as shown below:

```json
[
  {
    "operation": "buy",
    "unit-cost": 10.00,
    "quantity": 10000
  },
  {
    "operation": "sell",
    "unit-cost": 20.00,
    "quantity": 5000
  },
  {
    "operation": "sell",
    "unit-cost": 5.00,
    "quantity": 5000
  }
]
```
### Running with a file
To run the Tax Calculator using an input file, use the following command:
```bash
tax-calculator --file transaction.json
```
Replace transaction.json with the actual input file containing your transaction data.

### Running with JSON Input Directly
To pass JSON data directly via the command line, use the following command:
```bash
tax-calculator --input '[{"operation":"buy","unit-cost":10.00,"quantity":100},{"operation":"sell","unit-cost":15.00,"quantity":50},{"operation":"sell","unit-cost":15.00,"quantity":50}]'
```
Make sure the JSON string is correctly formatted and enclosed in single quotes to avoid issues with shell parsing.

