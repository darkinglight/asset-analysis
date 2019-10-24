package asset

import (
	"context"
	"encoding/json"
	"errors"
	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var errInvalidParam = errors.New("invalid paramters")

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

	findIncomeHandler := kithttp.NewServer(
		makeFindIncomeEndpoint(s),
		decodeFindIncomeRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/asset/v1/income", addIncomeHandler).Methods("POST")
	r.Handle("/asset/v1/income/{statement_id}", findIncomeHandler).Methods("GET")

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

func decodeFindIncomeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	idString, ok := vars["statement_id"]
	if !ok {
		return nil, errInvalidParam
	}
	id, err := strconv.Atoi(idString)
	if err != nil {
		return nil, errInvalidParam
	}
	return findIncomeRequest{StatementId: id}, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-i")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

type errorer interface {
	error() error
}
