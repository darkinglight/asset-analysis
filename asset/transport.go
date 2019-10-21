package asset

import (
	"context"
	"encoding/json"
	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

func MakeHandler(s Service, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeError),
	}

	addIncomeHandler := kithttp.NewServer(
		makeAddIncomeEndpoint(s),
		decodeAddIncomeRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/asset/v1/income", addIncomeHandler).Methods("POST")

	return r
}

func decodeAddIncomeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body struct {
		StatementId    int `json:"statement_id"`
		BusinessIncome int `json:"business_income"`
		BusinessCost   int `json:"business_cost"`
		GrossProfit    int `json:"gross_profit"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return addIncomeRequest{
		StatementId:    body.StatementId,
		BusinessIncome: body.BusinessIncome,
		BusinessCost:   body.BusinessCost,
		GrossProfit:    body.GrossProfit,
	}, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError()
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-i")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-i")
	switch err {
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
