package main

import "testing"

// TestVendingMachine general test suite for VendingMachine
func TestVendingMachine(t *testing.T) {
	cases := []struct {
		description string
		options map[Choice]Drink
		price int
		deposit int
		choice Choice
		drink Drink
	}{
		{
			"choiceless machine delivers nothing", 
			map[Choice]Drink{},
			0,
			0,
			Cola,
			None,
		},
		{
			"choosing Cola delivers Coke", 
			map[Choice]Drink{Cola: Coke},
			0,
			0,
			Cola,
			Coke,
		},
		{
			"choosing FizzyOrange delivers Fanta", 
			map[Choice]Drink{
				Cola: Coke,
				FizzyOrange: Fanta,
			},
			0,
			0,
			FizzyOrange,
			Fanta,
		},
		{
			"delivers nothing without money", 
			map[Choice]Drink{
				Cola: Coke,
				FizzyOrange: Fanta,
			},
			200,
			0,
			FizzyOrange,
			None,
		},
		{
			"delivers Coke when exact amount was deposited", 
			map[Choice]Drink{
				Cola: Coke,
				FizzyOrange: Fanta,
			},
			200,
			200,
			Cola,
			Coke,
		},
		{
			"delivers Fanta when deposit is more than enough", 
			map[Choice]Drink{
				Cola: Coke,
				FizzyOrange: Fanta,
			},
			150,
			200,
			FizzyOrange,
			Fanta,
		},

		{
			"delivers nothing when deposit is not enough", 
			map[Choice]Drink{
				Cola: Coke,
				FizzyOrange: Fanta,
			},
			150,
			100,
			FizzyOrange,
			None,
		},
	}

	for _, testCase := range cases {
		machine := InitVendingMachine(testCase.options, testCase.price)
		machine.Deposit(testCase.deposit)
		deliveredDrink := machine.Deliver(testCase.choice)
		if deliveredDrink != testCase.drink {
			t.Errorf("%v: VendingMachine(%v) == %v, want %v", testCase.description, testCase.choice, deliveredDrink, testCase.drink)
		}
	}
}
