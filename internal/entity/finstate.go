package entity

import "capital-gains/internal/utils"

type Finstate struct {
	Loss                    float64
	WeightedAverageUnitCost float64
	CurrentQuantity         int
}

func (f *Finstate) SetWeightedAverageUnitCost(value float64) {
	f.WeightedAverageUnitCost = utils.RoundTo2Decimals(value)
}

func (f *Finstate) GetWeightedAverageUnitCost() float64 {
	return f.WeightedAverageUnitCost
}
