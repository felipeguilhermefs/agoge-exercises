package main

import "testing"

// TestVendingMachine general test suite for VendingMachine
func TestVendingMachine(t *testing.T) {
	cases := []struct {
		description string
		options map[Choice]Drink
		choice Choice
		drink Drink
	}{
		{
			"choiceless machine delivers nothing", 
			make(map[Choice]Drink),
			Cola,
			None,
		},
		{
			"choosing Cola delivers Coke", 
			map[Choice]Drink{Cola: Coke},
			Cola,
			Coke,
		},
		{
			"choosing FizzyOrange delivers Fanta", 
			map[Choice]Drink{
				Cola: Coke,
				FizzyOrange: Fanta,
			},
			FizzyOrange,
			Fanta,
		},
	}

	for _, testCase := range cases {
		machine := InitVendingMachine(testCase.options)
		deliveredDrink := machine.Deliver(testCase.choice)
		if deliveredDrink != testCase.drink {
			t.Errorf("VendingMachine(%v) == %v, want %v", testCase.choice, deliveredDrink, testCase.drink)
		}
	}
}
