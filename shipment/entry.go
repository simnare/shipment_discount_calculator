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
