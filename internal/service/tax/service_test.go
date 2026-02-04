package tax

import (
	"capital-gains-api/internal/entity"
	"testing"
)

func newFinstate() *entity.Finstate {
	finstate := &entity.Finstate{
		Loss: 0,
	}
	finstate.SetWeightedAverageUnitCost(0)
	finstate.SetCurrentQuantity(0)

	return finstate
}

func TestOperationTaxResult(t *testing.T) {

	t.Run("Buy operation", func(t *testing.T) {
		service := NewService()
		operation := &entity.Operation{
			Operation: entity.Buy,
			UnitCost:  10,
			Quantity:  5000,
		}
		var expectedTax float64 = 0
		tax, err := service.OperationTaxResult(operation, newFinstate())
		if err != nil || tax != expectedTax {
			t.Errorf("Buy OperationTaxResult() = %v, %v; want %v, <nil>", tax, err, expectedTax)
		}
	})

	t.Run("Sell operation", func(t *testing.T) {
		service := NewService()
		operation := &entity.Operation{
			Operation: entity.Sell,
			UnitCost:  10,
			Quantity:  5000,
		}

		var expectedTax float64 = 10000
		tax, err := service.OperationTaxResult(operation, newFinstate())
		if err != nil || tax != expectedTax {
			t.Errorf("Sell OperationTaxResult() = %v, %v; want %v, <nil>", tax, err, expectedTax)
		}
	})

	t.Run("Unknown operation", func(t *testing.T) {
		service := NewService()
		operation := &entity.Operation{
			Operation: "Unknown",
		}
		_, err := service.OperationTaxResult(operation, newFinstate())
		if err == nil {
			t.Error("OperationTaxResult() with unknown operation; want error, got <nil>")
		}
	})
}
