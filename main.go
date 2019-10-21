package main

import (
	"asset-analysis/inmem"
	"context"
	"fmt"
	"github.com/go-kit/kit/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var (
		incomes = inmem.NewIncomeRepository()
		ctx     = context.Background()
	)
	httpAddr := ":8080"
	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	var as asset.Service
	as = assset.NewService(incomes)

	mux := http.NewServeMux()
	mux.Handle("/asset/v1/", asset.MakeHandler(as, logger))

	errs := make(chan error, 2)
	go func() {
		logger.Log("transport", "http", "address", httpAddr, "msg", "listening")
		errs <- http.ListenAndServe(httpAddr, nil)
	}()
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()
	logger.Log("terminated", errs)
}
