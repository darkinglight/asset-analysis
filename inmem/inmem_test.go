package inmem

import (
	"asset-analysis/income"
	"testing"
)

func TestIncomeSave(t *testing.T) {
	incomes := NewIncomeRepository()
	data := &income.Income{
		StatementId:    1,
		BusinessIncome: 1,
		BusinessCost:   1,
		GrossProfit:    1,
	}
	err := incomes.Save(data)
	if err != nil {
		t.Error("failed")
	}
}
