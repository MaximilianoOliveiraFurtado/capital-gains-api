package entity

import "capital-gains-api/internal/mathutil"

type Finstate struct {
	Loss                    float64
	weightedAverageUnitCost float64
	currentQuantity         int
}

func (f *Finstate) SetWeightedAverageUnitCost(value float64) {
	f.weightedAverageUnitCost = mathutil.RoundTo2Decimals(value)
}

func (f *Finstate) GetWeightedAverageUnitCost() float64 {
	return f.weightedAverageUnitCost
}

func (f *Finstate) IncrementQuantityCurrentQuantity(value int) {
	f.currentQuantity += value
}

func (f *Finstate) DecrementCurrentQuantity(value int) {

	f.currentQuantity -= value
	if f.currentQuantity == 0 {
		f.SetWeightedAverageUnitCost(0)
	}

}

func (f *Finstate) GetCurrentQuantity() int {
	return f.currentQuantity
}

func (f *Finstate) SetCurrentQuantity(value int) {
	f.currentQuantity = value
}
