package shipment

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func NewEntry(props ...string) Entry {
	e := Entry{}
	e.OriginalEntry = props

	if len(props) < 2 {
		log.Fatal("Shipment entry MUST have at least 2 params")
		return e
	}

	e.Date, _ = time.Parse("2006-01-02", props[0])
	e.Provider = Provider(props[1])

	if len(props) == 3 {
		e.Size = Size(props[2])
	}

	return e
}

func (e Entry) AsString() string {
	entries := e.OriginalEntry

	if e.Price != nil {
		entries = append(entries, fmt.Sprintf("%0.2d", e.Price))

		if e.DiscountPrice != nil {
			entries = append(entries, fmt.Sprintf("%0.2d", e.DiscountPrice))
		} else {
			entries = append(entries, "-")
		}
	} else {
		entries = append(entries, "Ignored")
	}

	return strings.Join(entries, " ")
}
