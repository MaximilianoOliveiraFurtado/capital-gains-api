package operation

import (
	"fmt"
	"os"

	"capital-gains/internal/entity"
	"capital-gains/internal/service/tax"
)

type IService interface {
	OperationTax(operation *entity.Operation) *entity.Tax
	InputParseOperation(operationsInputed string) []entity.Operation
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

func (s *Service) OperationTax(operation *entity.Operation) *entity.Tax {

	taxEntity, err := s.taxService.OperationTaxResult(operation)
	if err != nil {
		fmt.Fprintln(os.Stdout, []any{"unexpected error %s: {%s}", operation, err}...)
	}
	return &entity.Tax{
		Tax: taxEntity,
	}
}

func (s *Service) InputParseOperation(operationsInputed string) []entity.Operation {

	operation, err := entity.ParseOperations(operationsInputed)
	if err != nil {
		fmt.Fprintln(os.Stdout, []any{"error reading operation %s: {%s}", operation, err}...)
	}
	return operation

}
