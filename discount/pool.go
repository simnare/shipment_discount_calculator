package discount

import (
	"github.com/simnare/shipdisc/shipment"
)

type Pool struct {
	Rules       []Rule
	CreditGuard CreditLimitGuard
}

// NewPool initializes new rules pool
func NewPool(g CreditLimitGuard) *Pool {
	return &Pool{
		Rules:       make([]Rule, 0),
		CreditGuard: g}
}

// Add appends new discount rule to the rules pool
func (p *Pool) Add(r Rule) {
	p.Rules = append(p.Rules, r)
}

// CalculateDiscount returns discount object with applied discount sum and resulting price.
// Returns first applicable discount. Possible optimisation - apply all discounts.
func (p *Pool) CalculateDiscount(entry shipment.Entry) *Discount {
	dis := getApplicableDiscount(p.Rules, entry)

	if dis == nil {
		return dis
	}

	return p.CreditGuard.Check(entry, dis)
}

func getApplicableDiscount(rules []Rule, entry shipment.Entry) *Discount {
	for _, r := range rules {
		d := r.Apply(entry)
		if d != nil {
			return d
		}
	}

	return nil
}
