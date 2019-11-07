package discount

import (
	"math"

	"github.com/simnare/shipdisc/shipment"
)

type PackageSizeMinimumPrice struct {
	Size  shipment.Size
	Price float32
}

func NewPackageSizeMinimumPrice(size shipment.Size, pricesConfig shipment.PricesConfig) PackageSizeMinimumPrice {

	var minPrice float32 = math.MaxFloat32
	for _, packages := range pricesConfig.Prices {
		if price, ok := packages[size]; ok {
			minPrice = float32(math.Min(float64(minPrice), float64(price)))
		}
	}

	return PackageSizeMinimumPrice{Size: size, Price: minPrice}

}

func (r PackageSizeMinimumPrice) Apply(e shipment.Entry) *Discount {
	if e.Price == nil {
		return nil
	}

	ePrice := *e.Price
	if e.Size == r.Size && r.Price != ePrice {
		return &Discount{Price: r.Price, Discount: ePrice - r.Price}
	}

	return nil
}
