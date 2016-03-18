package main

import "testing"

func TestSubtotalCalculation(t *testing.T) {
	var units int
	var price float64
	var subtotal float64

	subtotal = GetSubtotal(units, price)
	if subtotal != 0 {
		t.Log("Sanity check, zero price check doesn't work")
		t.FailNow()
	}
}

func TestSubtotalUnitCalculation(t *testing.T) {
	var units int = 1
	var price float64 = 10.0
	var subtotal float64

	subtotal = GetSubtotal(units, price)
	if subtotal != 10.0 {
		t.Log("not counting units")
		t.FailNow()
	}
}

func TestTaxCalculation(t *testing.T) {
	var subtotal float64 = 100.0
	const taxrate = 8.0

	taxed := ApplyTax(subtotal, taxrate)

	if taxed != 108.0 {
		t.Log("Tax rate %f, should be 108.0")
		t.Fail()
	}

}

func TestGetTaxRate(t *testing.T) {
	var rate float64
	rate = GetTaxRate("UT")
	if rate != 6.85 {
		t.Log("Tax rate %f, for UT should be 6.85")
		t.Fail()
	}
	rate = GetTaxRate("CA")
	if rate != 8.4 {
		t.Log("Tax rate %f, for CA should be 8.2")
		t.Fail()
	}
}

func TestApplyDiscount(t *testing.T) {
	var value, subtotal float64 = 100.0, 0.0
	subtotal = ApplyDiscount(value)
	if subtotal != 97.0 {
		t.Log("discounted value %f, for input %f should be 97.0", subtotal, value)
		t.Fail()
	}
}
