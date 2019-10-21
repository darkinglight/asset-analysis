package asset

import (
	"context"
	"encoding/json"
	"net/http"
)

func decodeAddIncomeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body struct {
		StatementId    int `json:"statement_id"`
		BusinessIncome int `json:"business_income"`
		BusinessCost   int `json:"business_cost"`
		GrossProfit    int `json:"gross_profit"`
	}
	if err := json.NewDecoder(addIncomeRequest).Decode(&body); err != nil {
		return nil, err
	}

	return addIncomeRequest{
		StatementId:    body.StatementId,
		BusinessIncome: body.BusinessIncome,
		BusinessCost:   body.BusinessCost,
		GrossProfit:    body.GrossProfit,
	}, nil
}
