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

func makeAddIncomeEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) error {
		req := request.(addIncomeRequest)
		err := s.AddIncome(
			req.StatementId,
			req.BusinessIncome,
			req.BusinessCost,
			req.GrossProfit,
		)
		return err
	}
}

type FindIncomeRequest struct {
	StatementId int
}

type FindIncomeResponse struct {
	StatementId    int `json:"statement_id,omitempty"`
	BusinessIncome int `json:"business_income,omitempty"`
	BusinessCost   int `json:"business_cost,omitempty"`
	GrossProfit    int `json:"gross_profit,omitempty"`
}

func makeFindIncomeEndpoint(s Service) endpoint.Endpoint {
	return func(ctx contxt.Context, request interface{}) (interface{}, error) {
		req := request.(FindIncomeRequest)
		data, err := s.FindIncome(req.StatementId)
		return data, err
	}
}
