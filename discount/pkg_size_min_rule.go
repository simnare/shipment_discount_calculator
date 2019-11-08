package discount

import (
	"math"

	"github.com/simnare/shipdisc/shipment"
)

// PackageSizeMinimumPrice is a rule which applies minimal price from all providers for certain package size.
type PackageSizeMinimumPrice struct {
	Size  shipment.Size
	Price float32
}

// NewPackageSizeMinimumPrice initializes rule instance and calculates minimum price for given package size from all providers.
func NewPackageSizeMinimumPrice(size shipment.Size, pricesConfig shipment.PricesConfig) PackageSizeMinimumPrice {

	var minPrice float32 = math.MaxFloat32
	for _, packages := range pricesConfig.Prices {
		if price, ok := packages[size]; ok {
			minPrice = float32(math.Min(float64(minPrice), float64(price)))
		}
	}

	return PackageSizeMinimumPrice{Size: size, Price: minPrice}

}

// Apply checks if given shipment entry size matches the rule and returns discount with minimal price set.
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
