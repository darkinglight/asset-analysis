package asset

import (
	"asset-analysis/income"
	"asset-analysis/util"
	"encoding/json"
)

type Service interface {
	AddIncome(statementId int, businessIncome int, businessCost int, grossProfit int) error

	FindIncome(statementId int) *Income

	StoreIncome() error

	LoadIncome() error
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

func (s service) FindIncome(statementId int) *Income {
	income := s.income.Find(statementId)
	if income == nil {
		return nil
	}
	return &Income{
		StatementId:    income.StatementId,
		BusinessIncome: income.BusinessIncome,
		BusinessCost:   income.BusinessCost,
		GrossProfit:    income.GrossProfit,
	}
}

func (s service) StoreIncome() error {
	incomes := s.income.FindAll()
	storage, err := util.NewStore(util.IncomeStorage)
	if err != nil {
		return err
	}
	storage.Write(incomes)
	storage.Close()
	return nil
}

func (s service) LoadIncome() error {
	storage, err := util.NewStore(util.IncomeStorage)
	if err != nil {
		return err
	}
	data, err := storage.Read()
	var incomes []*income.Income
	json.Unmarshal(data, &incomes)

	err = s.income.SaveAll(incomes)
	if err != nil {
		return err
	}

	return nil
}

func NewService(income income.Repository) Service {
	return &service{
		income: income,
	}
}
