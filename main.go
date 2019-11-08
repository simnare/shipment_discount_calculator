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
	shippingPrices.Set("S", "LP", 1.5)
	shippingPrices.Set("M", "LP", 4.9)
	shippingPrices.Set("L", "LP", 6.9)
	shippingPrices.Set("S", "MR", 2)
	shippingPrices.Set("M", "MR", 3)
	shippingPrices.Set("L", "MR", 4)

	pool := discount.NewPool(discount.NewCreditLimitGuard(10))
	pool.Add(discount.NewPackageSizeMinimumPrice("S", *shippingPrices))
	pool.Add(discount.NewNthPackageFreePerMonth("LP", "L", 3))

	for sc.Scan() {
		entry := shipment.NewEntry(strings.Fields(sc.Text())...)
		entry.Price = shippingPrices.Get(entry.Provider, entry.Size)

		discount := pool.CalculateDiscount(entry)

		if discount != nil {
			entry.Price = &discount.Price
			entry.DiscountPrice = &discount.Discount
		}

		fmt.Println(entry)
	}

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
}
