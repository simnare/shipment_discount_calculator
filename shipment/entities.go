package shipment

import (
	"time"
)

// Providers define all known system shipment providers
var Providers = [2]Provider{"MR", "LP"}

// Sizes define all known system shipment sizes
var Sizes = [3]Size{"S", "M", "L"}

// Provider defines provider title type
type Provider string

// Size defines shipment size type
type Size string

type Entry struct {
	Date          time.Time
	Provider      Provider
	Size          Size
	Price         *float32
	DiscountPrice *float32
}

type PriceConfig struct {
	Provider Provider
	Size     Size
	Price    float32
}

type PricesConfig struct {
	Prices map[Provider]map[Size]float32
}

func NewPricesConfig() *PricesConfig {
	return &PricesConfig{Prices: make(map[Provider]map[Size]float32)}

}

func (r *PricesConfig) Set(c PriceConfig) {
	if _, ok := r.Prices[c.Provider]; !ok {
		r.Prices[c.Provider] = make(map[Size]float32, 0)
	}

	r.Prices[c.Provider][c.Size] = c.Price

	return
}

func isValidProvider(provider Provider) bool {
	switch provider {
	case "LM", "LP":
		return true
	}

	return false
}

func isValidSize(size Size) bool {
	switch size {
	case "S", "M", "L":
		return true
	}

	return false
}
