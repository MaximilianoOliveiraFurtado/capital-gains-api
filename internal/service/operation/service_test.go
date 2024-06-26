package operation

import (
	"capital-gains-api/internal/entity"
	"reflect"
	"testing"

	"capital-gains-api/internal/service/tax"

	"github.com/stretchr/testify/mock"
)

type MockTaxService struct {
	mock.Mock
}

func (m *MockTaxService) OperationTaxResult(operation *entity.Operation) (float64, error) {
	args := m.Called(operation)
	return args.Get(0).(float64), args.Error(1)
}

func newFinstate() *entity.Finstate {
	finstate := &entity.Finstate{
		Loss: 0,
	}
	finstate.SetWeightedAverageUnitCost(0)
	finstate.SetCurrentQuantity(0)

	return finstate
}

func TestOperationTaxSuccess(t *testing.T) {
	mockTaxService := new(MockTaxService)
	service := &Service{
		taxService: mockTaxService,
	}

	operation := &entity.Operation{
		Operation: entity.Sell,
		UnitCost:  10,
		Quantity:  5000,
	}

	var expectedTax float64 = 10000
	mockTaxService.On("OperationTaxResult", operation).Return(expectedTax, nil)

	taxResult := service.OperationTax(operation)
	if taxResult.Tax != expectedTax {
		t.Errorf("expected %v, found %v", expectedTax, taxResult.Tax)
	}

	mockTaxService.AssertExpectations(t)
}

func TestInputParseOperation(t *testing.T) {
	service := NewService(tax.NewService(newFinstate()))

	input := `[{"operation":"sell", "unit-cost":10.00, "quantity": 5000}]`
	expectedOperations := []entity.Operation{}
	operation := entity.Operation{
		Operation: entity.Sell,
		UnitCost:  10,
		Quantity:  5000,
	}
	expectedOperations = append(expectedOperations, operation)

	operations := service.InputParseOperation(input)

	if !reflect.DeepEqual(operations, expectedOperations) {
		t.Errorf("expected %v, found %v", expectedOperations, operations)
	}
}
