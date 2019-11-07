package discount

import (
	"fmt"

	"github.com/simnare/shipdisc/shipment"
)

type CreditLimitGuard struct {
	Limit       float32
	MonthlyUsed map[string]float32
}

func NewCreditLimitGuard(limit float32) CreditLimitGuard {
	return CreditLimitGuard{
		Limit:       limit,
		MonthlyUsed: make(map[string]float32, 0)}
}

func (g CreditLimitGuard) Check(entry shipment.Entry, discount *Discount) *Discount {
	if discount == nil {
		return nil
	}

	year, month, _ := entry.Date.Date()
	key := fmt.Sprintf("%d%d", year, month)
	discountVal := discount.Discount

	g.MonthlyUsed[key] += discountVal

	if g.MonthlyUsed[key] > g.Limit {
		overBudget := g.MonthlyUsed[key] - g.Limit
		return &Discount{
			Price:    discount.Price + overBudget,
			Discount: discount.Discount - overBudget}

	}

	return discount

}
