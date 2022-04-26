package main

import "testing"

// TestVendingMachine general test suite for VendingMachine
func TestVendingMachine(t *testing.T) {
	cases := []struct {
		description string
		choices []Choice
		drink Drink
	}{
		{"choiceless machine delivers nothing", make([]Choice, 0), None},
		{"choosing Cola delivers Coke", []Choice{Cola}, Coke},
		{"choosing FizzyOrange delivers Fanta", []Choice{Cola, FizzyOrange}, Fanta},
	}

	for _, testCase := range cases {
		machine := InitVendingMachine(testCase.choices)
		deliveredDrink := machine.Deliver(Cola)
		if deliveredDrink != testCase.drink {
			t.Errorf("VendingMachine(%v) == %v, want %v", testCase.choices, deliveredDrink, testCase.drink)
		}
	}
}
