package tax

import (
	"capital-gains/internal/entity"
	"capital-gains/internal/utils"
)

const (
	taxRate                        = 20
	taxFreeThresholdOperationValue = 20000
)

type IService interface {
	OperationTaxResult(operation entity.Operation) float64
}

type Service struct {
	finstate *entity.Finstate
}

func NewService(finstate *entity.Finstate) IService {
	return &Service{
		finstate: finstate,
	}
}

func (s *Service) OperationTaxResult(operation entity.Operation) float64 {

	switch operation.Operation {

	case entity.Buy:
		return s.buyOperationTaxResult(operation.UnitCost, operation.Quantity)
	case entity.Sell:
		return s.sellOperationTaxResult(operation.UnitCost, operation.Quantity)
	default:
		return 0

	}
}

func (s *Service) sellOperationTaxResult(operationUnitCost float64, operationQuantity int) float64 {

	var weightedAverageUnitCost float64 = s.finstate.GetWeightedAverageUnitCost()

	s.finstate.CurrentQuantity -= operationQuantity
	if s.finstate.CurrentQuantity == 0 {
		s.finstate.SetWeightedAverageUnitCost(0)
	}

	var weightedAverageTotalCost float64 = weightedAverageUnitCost * float64(operationQuantity)
	var operationTotalCost float64 = operationUnitCost * float64(operationQuantity)
	var tax float64 = 0

	if operationTotalCost < weightedAverageTotalCost {
		loss := weightedAverageTotalCost - operationTotalCost
		s.finstate.Loss += loss
		return tax
	}

	gain := operationTotalCost - weightedAverageTotalCost
	taxlableValue := s.taxDeduction(gain)

	if s.taxExemption(operationTotalCost) {
		return tax
	}

	if gain > 0 {

		tax = s.taxDue(taxlableValue)

	}

	return tax

}

func (s *Service) buyOperationTaxResult(operationUnitCost float64, operationQuantity int) float64 {

	s.weightedAverageUnitCost(operationUnitCost, operationQuantity)
	s.finstate.CurrentQuantity += operationQuantity

	return 0

}

func (s *Service) weightedAverageUnitCost(operationUnitCost float64, operationQuantity int) {

	currentQuantity := s.finstate.CurrentQuantity
	currentWeightedAverage := s.finstate.GetWeightedAverageUnitCost()

	var newWeightedAverageUnitCost float64 = ((float64(currentQuantity) * currentWeightedAverage) +
		(float64(operationQuantity) * operationUnitCost)) /
		(float64(currentQuantity) + float64(operationQuantity))

	s.finstate.SetWeightedAverageUnitCost(newWeightedAverageUnitCost)

}

func (s *Service) taxExemption(operationTotalCost float64) bool {
	return operationTotalCost <= taxFreeThresholdOperationValue
}

func (s *Service) taxDeduction(operationTotalCost float64) float64 {
	if operationTotalCost <= s.finstate.Loss {
		loss := s.finstate.Loss - operationTotalCost
		s.finstate.Loss = loss
		return 0
	}
	taxlableValue := operationTotalCost - s.finstate.Loss
	s.finstate.Loss = 0
	return taxlableValue
}

func (s *Service) taxDue(taxlableValue float64) float64 {
	return utils.RoundTo2Decimals((taxlableValue * taxRate) / 100)
}
