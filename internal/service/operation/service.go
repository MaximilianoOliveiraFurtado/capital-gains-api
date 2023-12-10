package operation

import (
	"capital-gains/internal/entity"
	"capital-gains/internal/service/tax"
)

type IService interface {
	OperationInput(operation entity.Operation) entity.Tax
}

type Service struct {
	taxService tax.IService
}

func NewService() IService {
	taxService := tax.NewService(&entity.Finstate{})
	return &Service{
		taxService: taxService,
	}
}

func (s *Service) OperationInput(operation entity.Operation) entity.Tax {

	var taxEntity entity.Tax
	taxEntity.Tax = s.taxService.OperationTaxResult(operation)
	return taxEntity
}
