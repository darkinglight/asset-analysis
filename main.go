package main

import (
	"context"
	"github.com/go-kit/kit/log"
	"os"
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
	as = assset.NewService()
}
