package main

import (
	"flag"
	"fmt"
	"strings"
)

var taxRate map[string]float64 = map[string]float64{
	"UT": 6.85,
	"CA": 8.40,
	"TX": 6.25,
	"NV": 8.00,
	"AL": 4.00,
}

type discountCutoff struct {
	rate   float64
	cutoff float64
}

var discounts []discountCutoff = []discountCutoff{
	{rate: 3.0, cutoff: 1000.0},
	{rate: 5.0, cutoff: 5000.0},
	{rate: 7.0, cutoff: 7000.0},
	{rate: 10.0, cutoff: 10000.0},
	{rate: 15.0, cutoff: 15000.0},
}

func GetSubtotal(units int, price float64) (subtotal float64) {
	return float64(units) * price
}

func ApplyTax(subtotal, tax float64) (total float64) {
	total = subtotal * (1 + tax/100.0)
	return
}

func GetTaxRate(state string) float64 {
	return taxRate[strings.ToUpper(state)]
}

func ApplyDiscount(subtotal float64) (discounted float64) {
	discounted = subtotal
	for _, v := range discounts {
		if subtotal >= v.cutoff {
			discounted = subtotal * (1.0 - v.rate/100.0)
		}
	}
	return
}

var units *int = flag.Int("units", 0, "Number of units to sell")
var price *float64 = flag.Float64("price", 0.0, "Unit price")
var state *string = flag.String("state", "", "State to Check")

func main() {
	var subtotal, total, tax float64
	flag.Parse()
	subtotal = GetSubtotal(*units, *price)
	subtotal = ApplyDiscount(subtotal)
	tax = GetTaxRate(*state)
	total = ApplyTax(subtotal, tax)
	fmt.Printf("%f\n", total)
}
