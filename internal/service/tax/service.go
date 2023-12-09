package tax

const (
	taxRate                        = 20
	taxFreeThresholdOperationValue = 20000
)

type OperationStats struct {
	Loss                    float64
	WeightedAverageUnitCost float64
	CurrentQuantity         int
}

type IService interface {
	weightedAverageUnitCost(operationUnitCost float64, operationQuantity int) float64
	sellOperationResult(operationUnitCost float64, operationQuantity int) float64
	//taxDeduction()
	//taxExemption()
	//taxDue()
}

type Service struct {
	operationStats OperationStats
}

func NewService() IService {
	operationStats := OperationStats{
		Loss:                    0,
		WeightedAverageUnitCost: 0,
		CurrentQuantity:         0,
	}
	return &Service{
		operationStats: operationStats,
	}
}

func (s *Service) weightedAverageUnitCost(operationUnitCost float64, operationQuantity int) float64 {

	currentQuantity := s.operationStats.CurrentQuantity
	currentWeightedAverage := s.operationStats.WeightedAverageUnitCost

	return ((float64(currentQuantity) * currentWeightedAverage) +
		(float64(operationQuantity) * operationUnitCost)) /
		(float64(currentQuantity) + float64(operationUnitCost))

}

func (s *Service) sellOperationResult(operationUnitCost float64, operationQuantity int) float64 {

	var newWeightedAverageUnitCost float64 = s.weightedAverageUnitCost(operationUnitCost, operationQuantity)
	s.operationStats.WeightedAverageUnitCost = newWeightedAverageUnitCost
	var weightedAverageTotalCost float64 = newWeightedAverageUnitCost * float64(operationQuantity)
	var operationTotalCost float64 = operationUnitCost * float64(operationQuantity)
	var tax float64 = 0

	if operationTotalCost < weightedAverageTotalCost {
		// repository.setLossOperation
	}

	if operationTotalCost > weightedAverageTotalCost {
		// service.taxExemption
		// service.taxDeduction
		// tax = repository.taxDue
	}

	return tax

}
