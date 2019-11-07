package discount

import (
	"fmt"

	"github.com/simnare/shipdisc/shipment"
)

type NthPackageFreePerMonth struct {
	Provider shipment.Provider
	Size     shipment.Size
	Which    int

	Counter map[string]int
}

func NewNthPackageFreePerMonth(provider shipment.Provider, size shipment.Size, which int) NthPackageFreePerMonth {
	r := NthPackageFreePerMonth{Provider: provider, Size: size, Which: which}
	r.Counter = make(map[string]int, 0)
	return r
}

func (r NthPackageFreePerMonth) Apply(e shipment.Entry) *Discount {
	if e.Price == nil {
		return nil
	}

	if e.Provider != r.Provider || e.Size != r.Size {
		return nil
	}

	year, month, _ := e.Date.Date()
	key := fmt.Sprintf("%d%d", year, month)
	if _, ok := r.Counter[key]; !ok {
		r.Counter[key] = 1
		return nil
	}

	r.Counter[key]++
	if r.Counter[key] == r.Which {
		return &Discount{Price: 0, Discount: *e.Price}
	}

	return nil
}
