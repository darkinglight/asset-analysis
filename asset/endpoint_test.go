package asset

import (
	"asset-analysis/inmem"
	"context"
	"testing"
)

func TestAddIncome(t *testing.T) {
	incomes := inmem.NewIncomeRepository()
	service := NewService(incomes)
	endpoint := makeAddIncomeEndpoint(service)
	request := addIncomeRequest{
		StatementId:    1,
		BusinessIncome: 2,
		BusinessCost:   3,
		GrossProfit:    4,
	}
	ctx := context.TODO()
	response, err := endpoint(ctx, request)
	if err != nil {
		t.Error("endpoint(AddIncome) add fail")
	}
	if response.(*addIncomeResponse).Err != nil {
		t.Error("endpoint(AddIncome) add fail")
	}
}
