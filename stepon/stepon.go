package stepon

import (
	"fmt"
	"sort"
	"strings"
)

var taxRate map[string]float64 = map[string]float64{
	"AL": 4.00,
	"CA": 8.40,
	"NV": 8.00,
	"TX": 6.25,
	"UT": 6.85,
}

type discountCutOff struct {
	rate   float64
	cutoff float64
}

var discounts []discountCutOff = []discountCutOff{
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

func GetStates() string {
	var states []string = make([]string, 0)
	for k := range taxRate {
		states = append(states, k)
	}
	sort.Sort(sort.StringSlice(states))
	return strings.Join(states, ",")
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

func GetAppName() string {
	return "Stepon Androids!"
}

func GetTotal(units int, price float64, state string) string {
	subtotal := GetSubtotal(units, price)
	discounted := ApplyDiscount(subtotal)
	total := ApplyTax(discounted, GetTaxRate(state))
	return fmt.Sprintf("Total: %f", total)
}

// var units *int = flag.Int("units", 0, "Number of units to sell. Defaults to 0")
// var price *float64 = flag.Float64("price", 0.0, "Unit price, defaults to 0.0.")
// var state *string = flag.String("state", "", "State to Check. If no state is included, assume no sales tax (out of state).")
