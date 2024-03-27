package main

import (
	"context"

	"capital-gains/internal/entity"
	"capital-gains/internal/service/operation"
)

type IController interface {
	PostTaxOperation(_ context.Context, []entity.Operation) ([]entity.Tax, error)
}

type operationTaxController struct {
	operationService operation.IService
}

func NewOperationTaxController(service operation.IService) IController {
	return &operationTaxController{
		operationService: operationService,
	}
}

func (c *operationTaxController) PostTaxOperation(ctx context.Context, operations []entity.Operation) ([]entity.Tax, error) {
	for _, operationInputed := range operations {
		operationTax := operationService.OperationTax(&operationInputed)
		operationsTaxes = append(operationsTaxes, *operationTax)
	}

	return operationsTaxes

}
