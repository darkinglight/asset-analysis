//income domain model
package income

import (
	"errors"
)

var ErrInvalidArgument = errors.New("invalid argument")

//利润表
type Income struct {
	//财务报表id
	StatementId int
	//营业收入
	BusinessIncome int
	//营业成本
	BusinessCost int
	//毛利润
	GrossProfit int
	//营业费用
	BusinessFee int
	//营业税费
	BusinessTax int
	//营业利润
	BusinessProfit int
	//外部收入
	OtherIncome int
	//外部支出
	OtherExpand int
	//税费
	Tax int
	//净利润
	NetProfit int
}

func (income Income) getGrossProfitRate() int {
	return income.GrossProfit / income.BusinessIncome
}

//利润表存储接口
type Repository interface {
	//save one income data
	Save(data Income) error

	//fetch one income data
	Find(id int) Income

	//find list income data
	FindAll() []Income
}
