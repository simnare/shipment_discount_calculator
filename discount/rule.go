package discount

import "github.com/simnare/shipdisc/shipment"

type Discount struct {
	Price    float32
	Discount float32
}

type Rule interface {
	Apply(shipment.Entry) *Discount
}
