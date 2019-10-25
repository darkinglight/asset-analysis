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

	data := service.FindIncome(1)
	if data.GrossProfit != 4 {
		t.Error("data not consist")
	}
}

func Test_StoreIncome(t *testing.T) {
	incomes := inmem.NewIncomeRepository()
	service := NewService(incomes)
	service.AddIncome(1, 2, 3, 4)
	err := service.StoreIncome()
	if err != nil {
		t.Error("store failed")
	}
}

func Test_LoadIncome(t *testing.T) {
	incomes := inmem.NewIncomeRepository()
	service := NewService(incomes)
	err := service.LoadIncome()
	if err != nil {
		t.Error("load failed")
	}
	item := service.FindIncome(1)
	if item == nil {
		t.Error("nothing returned")
	} else if item.GrossProfit != 4 {
		t.Errorf("GrossProfit expect 4, %d received", item.BusinessCost)
	}
}
