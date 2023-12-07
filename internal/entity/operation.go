package entity

type OperationType string

const (
	Buy  OperationType = "buy"
	Sell OperationType = "sell"
)

type Operation struct {
	Operation OperationType `json:"operation"`
	UnitCost  float64       `json:"unit_cost"`
	Quantity  int           `json:"quantity"`
}
