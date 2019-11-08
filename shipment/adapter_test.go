package shipment_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/simnare/shipdisc/shipment"
)

func TestNewEntry(t *testing.T) {
	dateString := "2015-01-29"
	date, _ := time.Parse("2006-01-02", dateString)

	cases := []struct {
		in   []string
		want shipment.Entry
	}{
		{
			[]string{dateString, "S", "LP"},
			shipment.Entry{
				OriginalEntry: []string{dateString, "S", "LP"},
				Date:          date,
				Provider:      shipment.Provider("LP"),
				Size:          shipment.Size("S")},
		},
		{
			[]string{dateString, "foo"},
			shipment.Entry{
				OriginalEntry: []string{dateString, "foo"},
				Date:          date,
				Size:          shipment.Size("foo")},
		},
	}

	for _, c := range cases {
		got := shipment.NewEntry(c.in...)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("NewEntry(%q) == %q, want = %q", c.in, got, c.want)
		}
	}
}

func TestString(t *testing.T) {
	dateString := "2015-01-29"
	date, _ := time.Parse("2006-01-02", dateString)

	cases := []struct {
		in   shipment.Entry
		want string
	}{
		{
			shipment.Entry{
				OriginalEntry: []string{dateString, "S", "LP"},
				Date:          date,
				Provider:      shipment.Provider("LP"),
				Size:          shipment.Size("S")},
			"2015-01-29 S LP Ignored",
		},
		{
			shipment.Entry{
				OriginalEntry: []string{dateString, "foo", "bar"},
				Date:          date,
				Provider:      shipment.Provider("bar"),
				Size:          shipment.Size("foo"),
				Price:         createPointerFloat(0.5)},

			"2015-01-29 foo bar 0.50 -",
		},
		{
			shipment.Entry{
				OriginalEntry: []string{dateString, "foo", "bar"},
				Date:          date,
				Provider:      shipment.Provider("bar"),
				Size:          shipment.Size("foo"),
				Price:         createPointerFloat(0.5),
				DiscountPrice: createPointerFloat(9.49)},
			"2015-01-29 foo bar 0.50 9.49",
		},
	}

	for _, c := range cases {
		got := c.in.String()
		if got != c.want {
			t.Errorf("Entry(%q).String() == %q, want = %q", c.in, got, c.want)
		}
	}
}

func createPointerFloat(v float32) *float32 {
	return &v
}
