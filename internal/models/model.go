package models

type Transaction struct {
	Operation string  `json:"operation"`
	UnitCost  float64 `json:"unit-cost"`
	Quantity  int     `json:"quantity"`
}

type TaxResult struct {
	Tax float64 `json:"tax"`
}
