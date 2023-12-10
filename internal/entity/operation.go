package entity

import (
	"encoding/json"
)

type OperationType string

const (
	Buy  OperationType = "buy"
	Sell OperationType = "sell"
)

type Operation struct {
	Operation OperationType `json:"operation"`
	UnitCost  float64       `json:"unit-cost"`
	Quantity  int           `json:"quantity"`
}

func ParseOperations(data string) ([]Operation, error) {
	var operations []Operation
	if err := json.Unmarshal([]byte(data), &operations); err != nil {
		return nil, err
	}
	return operations, nil
}
