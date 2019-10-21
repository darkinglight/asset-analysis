package asset

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type addIncomeRequest struct {
	StatementId    int
	BusinessIncome int
	BusinessCost   int
	GrossProfit    int
}

type addIncomeResponse struct {
	Err error `json:"error,omitempty"`
}

func (r addIncomeResponse) error() error { return r.Err }

func makeAddIncomeEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addIncomeRequest)
		err := s.AddIncome(
			req.StatementId,
			req.BusinessIncome,
			req.BusinessCost,
			req.GrossProfit,
		)
		return &addIncomeResponse{Err: err}, nil
	}
}

type findIncomeRequest struct {
	StatementId int
}

type Income struct {
	StatementId    int `json:"statement_id,omitempty"`
	BusinessIncome int `json:"business_income,omitempty"`
	BusinessCost   int `json:"business_cost,omitempty"`
	GrossProfit    int `json:"gross_profit,omitempty"`
}

type findIncomeResponse struct {
	Income *Income `json:"income,omitempty"`
	Err    error   `json:"error,omitempty"`
}

func (r findIncomeResponse) error() error { return r.Err }

func MakefindIncomeEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(findIncomeRequest)
		data, err := s.FindIncome(req.StatementId)
		return findIncomeResponse{Income: &data, Err: err}, nil
	}
}
