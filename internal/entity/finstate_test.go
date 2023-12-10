package entity

import (
	"testing"
)

func TestDecrementCurrentQuantity(t *testing.T) {
	finstate := Finstate{
		Loss:                    0,
		weightedAverageUnitCost: 10,
		currentQuantity:         5,
	}

	finstate.DecrementCurrentQuantity(3)
	if finstate.GetCurrentQuantity() != 2 {
		t.Errorf("DecrementCurrentQuantity() faill, expected %v, found %v", 2, finstate.GetCurrentQuantity())
	}
	if finstate.GetWeightedAverageUnitCost() != 10.0 {
		t.Errorf("DecrementCurrentQuantity() faill, expected %v, found %v", 10.0, finstate.GetWeightedAverageUnitCost())
	}

	finstate.DecrementCurrentQuantity(2)
	if finstate.GetCurrentQuantity() != 0 {
		t.Errorf("DecrementCurrentQuantity() faill, expected %v, found %v", 0, finstate.GetCurrentQuantity())
	}
	if finstate.GetWeightedAverageUnitCost() != 0 {
		t.Errorf("DecrementCurrentQuantity() faill, expected %v, found %v", 0.0, finstate.GetWeightedAverageUnitCost())
	}

}
