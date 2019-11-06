package shipment

import (
	"time"
)

// Provider defines provider title type
type Provider string

// Size defines shipment size type
type Size string

type Entry struct {
	OriginalEntry []string
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

func (r *PricesConfig) FindBasePrice(p Provider, s Size) *float32 {
	if v, ok := r.Prices[p][s]; ok {
		return &v
	}

	return nil
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
