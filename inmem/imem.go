package inmem

import (
	"sync"
)

type incomeRepository struct {
	mtx    sync.RWMutex
	autoId int
	incoms map[int]*income.Income
}

func (r *incomeRepository) Save(i *income.Income) error {
	r.mtx.Lock()
	defer r.mtx.UnLock()
	r.autoId++
	r.incoms[r.autoId] = i
	return nil
}
