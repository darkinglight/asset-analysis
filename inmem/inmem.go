package inmem

import (
	"asset-analysis/income"
	"sync"
)

type incomeRepository struct {
	mtx     sync.RWMutex
	incomes map[int]*income.Income
}

func (r *incomeRepository) Save(i *income.Income) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.incomes[i.StatementId] = i
	return nil
}

func (r *incomeRepository) Find(statementId int) *income.Income {
	result := r.incomes[statementId]
	return result
}

func (r *incomeRepository) SaveAll(data []*income.Income) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	for item := range data {
		r.incomes[item.StatementId] = item
	}
	return nil
}

func (r *incomeRepository) FindAll() []*income.Income {
	var result []*income.Income
	for _, income := range r.incomes {
		result = append(result, income)
	}
	return result
}

func NewIncomeRepository() income.Repository {
	return &incomeRepository{
		incomes: make(map[int]*income.Income),
	}
}
