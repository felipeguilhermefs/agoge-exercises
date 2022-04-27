package main

import "testing"

// TestChoicelessVendingMachine test behaviour when there is no choice
func TestChoicelessVendingMachine(t *testing.T) {
	noChoice := map[Choice]Drink{}
	machine := InitVendingMachine(noChoice, 0)

	testCases := []struct {
		choice Choice
		drink Drink
	}{
		{Cola, None},
		{FizzyOrange, None},
	}

	for _, testCase := range testCases {
		deliveredDrink := machine.Deliver(testCase.choice)
		if deliveredDrink != testCase.drink {
			t.Errorf("TestChoicelessVendingMachine (%v) == %v, want %v", testCase.choice, deliveredDrink, testCase.drink)
		}
	}
}

func TestFreeVendingMachine(t *testing.T) {
	options := map[Choice]Drink{
		Cola: Coke,
		FizzyOrange: Fanta,
	}
	noPrice := 0
	machine := InitVendingMachine(options, noPrice)

	testCases := []struct {
		choice Choice
		drink Drink
	}{
		{Cola, Coke},
		{FizzyOrange, Fanta},
	}

	for _, testCase := range testCases {
		deliveredDrink := machine.Deliver(testCase.choice)
		if deliveredDrink != testCase.drink {
			t.Errorf("TestChoicelessVendingMachine (%v) == %v, want %v", testCase.choice, deliveredDrink, testCase.drink)
		}
	}
}

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
