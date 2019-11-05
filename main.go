package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

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
	for sc.Scan() {
		fmt.Println(sc.Text())
	}

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	shippingPrices := shipment.NewPricesConfig()
	shippingPrices.Set(shipment.PriceConfig{Provider: "LP", Size: "S", Price: 1.5})
	shippingPrices.Set(shipment.PriceConfig{Provider: "LP", Size: "M", Price: 4.9})
	shippingPrices.Set(shipment.PriceConfig{Provider: "LP", Size: "L", Price: 6.9})
	shippingPrices.Set(shipment.PriceConfig{Provider: "MR", Size: "S", Price: 2})
	shippingPrices.Set(shipment.PriceConfig{Provider: "MR", Size: "M", Price: 3})
	shippingPrices.Set(shipment.PriceConfig{Provider: "MR", Size: "L", Price: 4})

}
