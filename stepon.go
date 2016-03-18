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

func ApplyDiscount(subtotal float64) float64 {
	return subtotal * (1.0 - 3.0/100.0)
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
