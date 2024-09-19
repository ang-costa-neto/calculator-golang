package models

/**
 * Declara a estrutura de dados para o transaction
 * aqui usa a tag `json:"field-name"` para definir como os campos serializados/deserializados
 * também pode usar a tag `xml:"field-name"` para mapear em xml
 * quanto a struct for convertido para JSON e vice-versa
 */
type Transaction struct {
	Code      string  `json:"code"`      // O campo Code será mapeado para o campo "code" no JSON
	Operation string  `json:"operation"` // O campo Operation será mapeado para o campo "operation" no JSON
	UnitCost  float64 `json:"unit-cost"` // O campo UnitCost será mapeado para o campo "unit-cost" no JSON
	Quantity  int     `json:"quantity"`  // O campo Quantity será mapeado para o campo "quantity" no JSON
}

/**
 * Declara a estrutua de dados para a tax result
 */
type TaxResult struct {
	Code string  `json:"code"` // O campo Code será mapeado para o campo "code" no JSON
	Tax  float64 `json:"tax"`  // O campo Taz será mapeado para o campo "tax" no JSON
}

/**
 * Declara a estrutura de dados para o operation
 */
type Operation struct {
	Quantity             int
	WeightedAveragePrice float64
	Prejudice            float64
}
