package handler

import (
	"errors"

	"github.com/ang-costa-neto/calculator-golang/internal/constants"
	"github.com/ang-costa-neto/calculator-golang/internal/models"
)

type Processor interface {
	ProcessTransactions(transactions []models.Transaction) ([]models.TaxResult, error)
}

type realProcessor struct{}

func calculateWeightedAveragePrice(currentQuantity int, weightedAveragePrice float64, quantityBought int, priceBought float64) float64 {
	totalQuantity := float64(currentQuantity + quantityBought)
	return ((float64(currentQuantity) * weightedAveragePrice) + (float64(quantityBought) * priceBought)) / totalQuantity
}

func (p *realProcessor) ProcessTransactions(transactions []models.Transaction) ([]models.TaxResult, error) {
	var taxes []models.TaxResult
	operations := make(map[string]models.Operation)

	for i := range transactions {
		code := transactions[i].Code
		operationType := transactions[i].Operation
		quantity := transactions[i].Quantity
		unitCost := transactions[i].UnitCost

		if _, ok := operations[code]; !ok {
			operations[code] = models.Operation{
				Quantity:             0,
				WeightedAveragePrice: 0.0,
				Prejudice:            0.0,
			}
		}

		op := operations[code]

		switch operationType {
		case constants.BuyOperation:
			op.WeightedAveragePrice = calculateWeightedAveragePrice(op.Quantity, op.WeightedAveragePrice, quantity, unitCost)
			op.Quantity += quantity
			taxes = append(taxes, models.TaxResult{Code: code, Tax: 0.0})
		case constants.SellOperation:
			capitalGain := (float64(quantity) * unitCost) - (float64(quantity) * op.WeightedAveragePrice) + op.Prejudice
			op.Prejudice = 0.0

			if unitCost < op.WeightedAveragePrice || capitalGain < 0 {
				op.Prejudice += capitalGain
				capitalGain = 0
			}

			if capitalGain > 0 && unitCost > op.WeightedAveragePrice && (float64(quantity)*unitCost) > 20000 {
				taxes = append(taxes, models.TaxResult{Code: code, Tax: capitalGain * constants.TAX})
			} else {
				taxes = append(taxes, models.TaxResult{Code: code, Tax: 0.0})
			}

			op.Quantity -= quantity
		default:
			return nil, errors.New("invalid operation")
		}

		operations[code] = op
	}

	return taxes, nil
}

var ProcessorInstance Processor = &realProcessor{}
