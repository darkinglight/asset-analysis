package main

import (
	"errors"
	"fmt"
)

var ErrInvalidArgument = errors.New("invalid argument")

type Service interface {
	//save one income data
	Save(data Income) error

	//fetch one income data
	Fetch(id int) Income
}

type Income struct {
	BusinessIncome int
	BusinessCost   int
	GrossProfit    int
	BusinessFee    int
	BusinessTax    int
	BusinessProfit int
	OtherIncome    int
	OtherExpand    int
	Tax            int
	NetProfit      int
}
