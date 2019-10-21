package asset

import (
	"fmt"
)

type Service interface {
	AddIncome(statementId int, businessIncome int, businessCost int, grossProfit int) error

	FindIncome(statementId int) *income.Income
}

type service struct {
	income income.incomeRepository
}

func (s service) AddIncome(statementId int, businessIncome int, businessCost int, grossProfit int) error {
	incomeData := &income.Income{
		StatementId:    statementId,
		BusinessIncome: businessIncome,
		BusinessCost:   businessCost,
		GrossProfit:    grossProfit,
	}
	saveError := s.income.Save(incomeData)
	return saveError
}
