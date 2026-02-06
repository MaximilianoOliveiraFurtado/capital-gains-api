package tax

import (
	"errors"

	"capital-gains-api/internal/domain/entity"
	"capital-gains-api/internal/util"
)

const (
	taxRate                        = 20
	taxFreeThresholdOperationValue = 20000
	UNKNOWN_OPERATION_ERROR        = "unknown operation"
)

type IService interface {
	OperationTaxResult(operation *entity.Operation, finstate *entity.Finstate) (float64, error)
}

type Service struct{}

func NewService() IService {
	return &Service{}
}

func (s *Service) OperationTaxResult(operation *entity.Operation, finstate *entity.Finstate) (float64, error) {

	switch operation.Operation {

	case entity.Buy:
		return s.buyOperationTaxResult(operation.UnitCost, operation.Quantity, finstate), nil
	case entity.Sell:
		return s.sellOperationTaxResult(operation.UnitCost, operation.Quantity, finstate), nil
	default:
		return -1, errors.New(UNKNOWN_OPERATION_ERROR)

	}
}

func (s *Service) sellOperationTaxResult(operationUnitCost float64, operationQuantity int, finstate *entity.Finstate) float64 {

	weightedAverageUnitCost := finstate.GetWeightedAverageUnitCost()

	finstate.DecrementCurrentQuantity(operationQuantity)
	weightedAverageTotalCost := weightedAverageUnitCost * float64(operationQuantity)
	var operationTotalCost float64 = operationUnitCost * float64(operationQuantity)
	var tax float64 = 0

	loss := s.lossCalculation(operationTotalCost, weightedAverageTotalCost, finstate)

	if loss > 0 {
		return tax
	}

	gain := operationTotalCost - weightedAverageTotalCost
	taxlableValue := s.taxDeduction(gain, finstate)

	if s.taxExemption(operationTotalCost) {
		return tax
	}

	if gain > 0 {

		tax = s.taxDue(taxlableValue)

	}

	return tax

}

func (s *Service) buyOperationTaxResult(operationUnitCost float64, operationQuantity int, finstate *entity.Finstate) float64 {

	s.weightedAverageUnitCost(operationUnitCost, operationQuantity, finstate)
	finstate.IncrementQuantityCurrentQuantity(operationQuantity)
	return 0

}

func (s *Service) weightedAverageUnitCost(operationUnitCost float64, operationQuantity int, finstate *entity.Finstate) {

	currentQuantity := finstate.GetCurrentQuantity()
	currentWeightedAverage := finstate.GetWeightedAverageUnitCost()

	var newWeightedAverageUnitCost float64 = ((float64(currentQuantity) * currentWeightedAverage) +
		(float64(operationQuantity) * operationUnitCost)) /
		(float64(currentQuantity) + float64(operationQuantity))

	finstate.SetWeightedAverageUnitCost(newWeightedAverageUnitCost)

}

func (s *Service) lossCalculation(operationTotalCost float64, weightedAverageTotalCost float64, finstate *entity.Finstate) float64 {
	if operationTotalCost < weightedAverageTotalCost {
		loss := weightedAverageTotalCost - operationTotalCost
		finstate.Loss += loss
		return loss
	}
	return 0
}

func (s *Service) taxExemption(operationTotalCost float64) bool {
	return operationTotalCost <= taxFreeThresholdOperationValue
}

func (s *Service) taxDeduction(operationTotalCost float64, finstate *entity.Finstate) float64 {
	if operationTotalCost <= finstate.Loss {
		loss := finstate.Loss - operationTotalCost
		finstate.Loss = loss
		return 0
	}
	taxlableValue := operationTotalCost - finstate.Loss
	finstate.Loss = 0
	return taxlableValue
}

func (s *Service) taxDue(taxlableValue float64) float64 {
	return util.RoundTo2Decimals((taxlableValue * taxRate) / 100)
}
