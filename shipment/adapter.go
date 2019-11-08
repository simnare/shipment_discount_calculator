package shipment

import (
	"fmt"
	"log"
	"strings"
	"time"
)

// NewEntry accepts slice of strings as fields to construct new shipment entry.
func NewEntry(props ...string) Entry {
	e := Entry{OriginalEntry: props}

	if len(props) < 2 {
		// TODO return error instead of fatal
		log.Fatal("Shipment entry MUST have at least 2 params")
	}

	e.Date, _ = time.Parse("2006-01-02", props[0])
	e.Size = Size(props[1])

	if len(props) == 3 {
		e.Provider = Provider(props[2])
	}

	return e
}

// String returns shipment entry as a string.
func (e Entry) String() string {
	entries := e.OriginalEntry

	if e.Price != nil {
		entries = append(entries, fmt.Sprintf("%0.2f", *e.Price))

		if e.DiscountPrice != nil {
			entries = append(entries, fmt.Sprintf("%0.2f", *e.DiscountPrice))
		} else {
			entries = append(entries, "-")
		}
	} else {
		entries = append(entries, "Ignored")
	}

	return strings.Join(entries, " ")
}
