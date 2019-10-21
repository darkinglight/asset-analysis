package asset

import (
	"asset-analysis/income"
)

type Service interface {
	AddIncome(statementId int, businessIncome int, businessCost int, grossProfit int) error

	FindIncome(statementId int) *income.Income
}

type service struct {
	income income.Repository
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

func (s service) FindIncome(statementId int) *income.Income {
	income := s.income.Find(statementId)
	return income
}

func NewService(income income.Repository) Service {
	return &service{
		income: income,
	}
}
