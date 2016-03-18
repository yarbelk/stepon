package main

import "testing"

func TestSubtotalCalculation(t *testing.T) {
	var units int
	var price float64
	var subtotal float64

	subtotal = GetSubtotal(units, price)
	if subtotal != 0 {
		t.Log("Sanity check, zero price check doesn't work\n")
		t.FailNow()
	}
}

func TestSubtotalUnitCalculation(t *testing.T) {
	var units int = 1
	var price float64 = 10.0
	var subtotal float64

	subtotal = GetSubtotal(units, price)
	if subtotal != 10.0 {
		t.Log("not counting units\n")
		t.FailNow()
	}
}

func TestTaxCalculation(t *testing.T) {
	var subtotal float64 = 100.0
	const taxrate = 8.0

	taxed := ApplyTax(subtotal, taxrate)

	if taxed != 108.0 {
		t.Logf("Tax rate %f, should be 108.0\n")
		t.Fail()
	}

}

func TestGetTaxRate(t *testing.T) {
	var rate float64
	rate = GetTaxRate("UT")
	if rate != 6.85 {
		t.Logf("Tax rate %f, for UT should be 6.85\n")
		t.Fail()
	}
	rate = GetTaxRate("CA")
	if rate != 8.4 {
		t.Logf("Tax rate %f, for CA should be 8.2\n")
		t.Fail()
	}
}

func TestApplyDiscounts(t *testing.T) {
	var subtotal float64

	var flagtests = []struct {
		in  float64
		out float64
	}{
		{100., 100.},
		{1000., 970.},
		{5000., 4750.},
		{10000., 9000.},
		{50000., 42500.},
	}

	for _, v := range flagtests {
		subtotal = ApplyDiscount(v.in)
		if subtotal != v.out {
			t.Logf("discounted value %f, for input %f should be 4851\n", subtotal, v.in)
			t.Fail()
		}
	}
}
