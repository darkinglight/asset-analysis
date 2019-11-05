package asset

import (
	"github.com/go-kit/kit/metrics"
	"time"
)

type instrumentingService struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	Service
}

func NewInstrumentingService(counter metrics.Counter, latency metrics.Histogram, s Service) Service {
	return &instrumentingService{
		requestCount:   counter,
		requestLatency: latency,
		Service:        s,
	}
}

func (s *instrumentingService) AddIncome(statementId int, businessIncome int, businessCost int, grossProfit int) error {
	defer func(begin time.Time) {
		s.requestCount.With("method", "addincome").Add(1)
		s.requestLatency.With("method", "addincome").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.AddIncome(statementId, businessIncome, businessCost, grossProfit)
}

func (s *instrumentingService) FindIncome(statementId int) *Income {
	defer func(begin time.Time) {
		s.requestCount.With("method", "findincome").Add(1)
		s.requestLatency.With("method", "findincome").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.FindIncome(statementId)
}

func (s *instrumentingService) StoreIncome() error {
	return s.Service.StoreIncome()
}

func (s *instrumentingService) LoadIncome() error {
	return s.Service.LoadIncome()
}
