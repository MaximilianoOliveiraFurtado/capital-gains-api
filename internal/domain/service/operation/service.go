package operation

import (
	"fmt"
	"os"

	"capital-gains-api/internal/domain/entity"
	"capital-gains-api/internal/domain/service/tax"
)

type IService interface {
	OperationTax(operation *entity.Operation, finstate *entity.Finstate) *entity.Tax
	InputParseOperation(operationsInputed string) []entity.Operation
}

type Service struct {
	taxService tax.IService
}

func NewService(taxService tax.IService) IService {
	return &Service{
		taxService: taxService,
	}
}

func (s *Service) OperationTax(operation *entity.Operation, finstate *entity.Finstate) *entity.Tax {

	tax, err := s.taxService.OperationTaxResult(operation, finstate)
	if err != nil {
		fmt.Fprintln(os.Stdout, []any{"unexpected error %s: {%s}", operation, err}...)
	}
	return &entity.Tax{
		Tax: tax,
	}
}

func (s *Service) InputParseOperation(operationsInputed string) []entity.Operation {

	operation, err := entity.ParseOperations(operationsInputed)
	if err != nil {
		fmt.Fprintln(os.Stdout, []any{"error reading operation %s: {%s}", operation, err}...)
	}
	return operation

}
