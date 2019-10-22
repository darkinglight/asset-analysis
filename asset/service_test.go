package asset

import (
	"asset-analysis/inmem"
	"testing"
)

func Test_AddIncome(t *testing.T) {
	incomes := inmem.NewIncomeRepository()
	service := NewService(incomes)
	err := service.AddIncome(1, 2, 3, 4)
	if err != nil {
		t.Error("save failed")
	}
}
