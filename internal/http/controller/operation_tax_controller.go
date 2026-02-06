package controller

import (
	"context"

	"capital-gains-api/internal/domain/entity"
	"capital-gains-api/internal/domain/service/operation"
)

type IController interface {
	PostTaxOperation(_ context.Context, _ []entity.Operation) ([]entity.Tax, error)
}

type operationTaxController struct {
	operationService operation.IService
}

func NewOperationTaxController(operationService operation.IService) IController {
	return &operationTaxController{
		operationService: operationService,
	}
}

func (c *operationTaxController) PostTaxOperation(ctx context.Context, operations []entity.Operation) ([]entity.Tax, error) {

	var operationsTaxes []entity.Tax
	finstate := &entity.Finstate{}
	for _, operationInputed := range operations {
		operationTax := c.operationService.OperationTax(&operationInputed, finstate)
		operationsTaxes = append(operationsTaxes, *operationTax)
	}

	return operationsTaxes, nil

}
