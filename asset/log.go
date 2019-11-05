package asset

import (
	"github.com/go-kit/kit/log"
	"time"
)

type loggingService struct {
	logger log.Logger
	Service
}

func NewLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{logger, s}
}

func (s *loggingService) AddIncome(statementId int, businessIncome int, businessCost int, grossProfit int) error {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "addincome",
			"statementId", statementId,
			"businessIncome", businessIncome,
		)
	}(time.Now())

	return s.Service.AddIncome(statementId, businessIncome, businessCost, grossProfit)
}

func (s *loggingService) FindIncome(statementId int) *Income {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "addincome",
			"statementId", statementId,
		)
	}(time.Now())

	return s.Service.FindIncome(statementId)
}

func (s *loggingService) StoreIncome() error {
	return s.Service.StoreIncome()
}

func (s *loggingService) LoadIncome() error {
	return s.Service.LoadIncome()
}
