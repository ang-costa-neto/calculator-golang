package handler

import (
	"errors"
	"tax-calculator/internal/constants"
	"tax-calculator/internal/models"
)

func calculateWeightedAveragePrice(
	currentQuantity int,
	weightedAveragePrice float64,
	quantityBought int,
	priceBought float64,
) float64 {
	totalQuantity := float64(currentQuantity + quantityBought)
	return ((float64(currentQuantity) * weightedAveragePrice) + (float64(quantityBought) * priceBought)) / totalQuantity
}

func ProcessTransactions(transactions []models.Transaction) ([]models.TaxResult, error) {
	var taxes []models.TaxResult
	currentQuantity := 0
	weightedAveragePrice := 0.0
	prejudice := 0.0

	for i := range transactions {
		switch transactions[i].Operation {
		case constants.BuyOperation:
			weightedAveragePrice = calculateWeightedAveragePrice(
				currentQuantity,
				weightedAveragePrice,
				transactions[i].Quantity,
				transactions[i].UnitCost,
			)
			currentQuantity += transactions[i].Quantity
			taxes = append(taxes, models.TaxResult{})
		case constants.SellOperation:
			capitalGain := (float64(transactions[i].Quantity) * transactions[i].UnitCost) - (float64(transactions[i].Quantity) * weightedAveragePrice) + prejudice
			prejudice = 0.0

			if transactions[i].UnitCost < weightedAveragePrice || capitalGain < 0 {
				prejudice += capitalGain
				capitalGain = 0
			}

			if capitalGain > 0 && transactions[i].UnitCost > weightedAveragePrice && (transactions[i].UnitCost*float64(transactions[i].Quantity)) > 20000 {
				taxes = append(taxes, models.TaxResult{Tax: capitalGain * constants.TAX})
			} else {
				taxes = append(taxes, models.TaxResult{})
			}

			currentQuantity -= transactions[i].Quantity
		default:
			return nil, errors.New("invalid operation")
		}
	}

	return taxes, nil
}
