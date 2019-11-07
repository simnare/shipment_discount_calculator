package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/simnare/shipdisc/discount"
	"github.com/simnare/shipdisc/shipment"
)

func main() {
	args := os.Args[1:]
	inputFile := "./input.txt"
	if len(args) == 1 {
		inputFile = args[0]
	}

	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	shippingPrices := shipment.NewPricesConfig()
	shippingPrices.Set(shipment.PriceConfig{Provider: "LP", Size: "S", Price: 1.5})
	shippingPrices.Set(shipment.PriceConfig{Provider: "LP", Size: "M", Price: 4.9})
	shippingPrices.Set(shipment.PriceConfig{Provider: "LP", Size: "L", Price: 6.9})
	shippingPrices.Set(shipment.PriceConfig{Provider: "MR", Size: "S", Price: 2})
	shippingPrices.Set(shipment.PriceConfig{Provider: "MR", Size: "M", Price: 3})
	shippingPrices.Set(shipment.PriceConfig{Provider: "MR", Size: "L", Price: 4})

	pool := discount.NewPool(discount.NewCreditLimitGuard(10))
	pool.Add(discount.NewPackageSizeMinimumPrice("S", *shippingPrices))
	pool.Add(discount.NewNthPackageFreePerMonth("LP", "L", 3))

	// var previous shipment.Entry
	for sc.Scan() {
		entry := shipment.NewEntry(strings.Fields(sc.Text())...)
		entry.Price = shippingPrices.FindBasePrice(entry.Provider, entry.Size)

		discount := pool.CalculateDiscount(entry)

		if discount != nil {
			entry.Price = &discount.Price
			entry.DiscountPrice = &discount.Discount
		}

		fmt.Println(entry.AsString())
		// previous = entry
	}

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
}
