package models

type Transaction struct {
	Code      string  `json:"code"`
	Operation string  `json:"operation"`
	UnitCost  float64 `json:"unit-cost"`
	Quantity  int     `json:"quantity"`
}

type TaxResult struct {
	Code string  `json:"code"`
	Tax  float64 `json:"tax"`
}

type Operation struct {
	Quantity             int
	WeightedAveragePrice float64
	Prejudice            float64
}
