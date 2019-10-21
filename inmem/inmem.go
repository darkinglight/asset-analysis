package inmem

import (
	"asset-analysis/income"
	"sync"
)

type incomeRepository struct {
	mtx    sync.RWMutex
	incoms map[int]*income.Income
}

func (r *incomeRepository) Save(i *income.Income) error {
	r.mtx.Lock()
	defer r.mtx.UnLock()
	r.incoms[i.StatementId] = i
	return nil
}

func (r *incomeRepository) Find(statementId int) *income.Income {
	result := r.incoms[statementId]
	return result
}

func (r *incomeRepository) FindAll() []*income.Income {
	return r.incoms
}

func NewIncomeRepository() income.Repository {
	return &incomeRepository{
		incomes: make(map[int]*income.Income),
	}
}
