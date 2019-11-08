package shipment

// PricesConfig constructs prices map.
type PricesConfig struct {
	Prices map[Provider]map[Size]float32
}

// NewPricesConfig initializes new config with empty values.
func NewPricesConfig() *PricesConfig {
	return &PricesConfig{Prices: make(map[Provider]map[Size]float32)}
}

// Set inserts price for provider and size.
func (r *PricesConfig) Set(size Size, provider Provider, price float32) {
	if _, ok := r.Prices[provider]; !ok {
		r.Prices[provider] = make(map[Size]float32, 0)
	}

	r.Prices[provider][size] = price
}

// Get returns price for given provider and size, nil otherwise.
func (r *PricesConfig) Get(provider Provider, size Size) *float32 {
	if v, ok := r.Prices[provider][size]; ok {
		return &v
	}

	return nil
}
