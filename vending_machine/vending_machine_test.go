package main

import "testing"

// TestVendingMachine general test suite for VendingMachine
func TestVendingMachine(t *testing.T) {
	cases := []struct {
		description string
		choices []string
		drink string
	}{
		{"choiceless machine delivers nothing", make([]string, 0), ""},
	}

	for _, testCase := range cases {
		machine := InitVendingMachine(testCase.choices)
		deliveredDrink := machine.Deliver("cola")
		if deliveredDrink != testCase.drink {
			t.Errorf("VendingMachine(%v) == %v, want %v", testCase.choices, deliveredDrink, testCase.drink)
		}
	}
}
