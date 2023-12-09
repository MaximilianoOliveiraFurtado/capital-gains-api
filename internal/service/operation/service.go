package operation

import (
	"capital-gains/internal/entity"
	"capital-gains/internal/service/tax"
)

type IService interface {
	OperationInput(data string) (*entity.Operation, error)
}

type Service struct {
	taxService tax.IService
}

func NewService(taxService tax.IService) IService {
	return &Service{
		taxService: taxService,
	}
}

func (s *Service) OperationInput(data string) (*entity.Operation, error) {

}
