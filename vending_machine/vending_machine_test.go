package main

import "testing"

// TestChoicelessVendingMachine test behaviour when there is no choice
func TestChoicelessVendingMachine(t *testing.T) {
	noChoice := map[Choice]Drink{}
	machine := InitVendingMachine(noChoice, 0)

	testCases := []struct {
		choice Choice
		drink  Drink
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

// TestFreeVendingMachine test behaviour when there is no price
func TestFreeVendingMachine(t *testing.T) {
	options := map[Choice]Drink{
		Cola:        Coke,
		FizzyOrange: Fanta,
	}
	noPrice := 0
	machine := InitVendingMachine(options, noPrice)

	testCases := []struct {
		choice Choice
		drink  Drink
	}{
		{Cola, Coke},
		{FizzyOrange, Fanta},
	}

	for _, testCase := range testCases {
		deliveredDrink := machine.Deliver(testCase.choice)
		if deliveredDrink != testCase.drink {
			t.Errorf("TestFreeVendingMachine (%v) == %v, want %v", testCase.choice, deliveredDrink, testCase.drink)
		}
	}
}

// TestPricedVendingMachine test behaviour when there is a price for drinks
func TestPricedVendingMachine(t *testing.T) {
	options := map[Choice]Drink{
		Cola:        Coke,
		FizzyOrange: Fanta,
	}
	price := 200

	testCases := []struct {
		name    string
		deposit int
		choice  Choice
		drink   Drink
	}{
		{
			"delivers nothing without money",
			0,
			FizzyOrange,
			None,
		},
		{
			"delivers Coke when exact amount was deposited",
			200,
			Cola,
			Coke,
		},
		{
			"delivers Fanta when deposit is more than enough",
			250,
			FizzyOrange,
			Fanta,
		},
		{
			"delivers nothing when deposit is not enough",
			150,
			FizzyOrange,
			None,
		},
	}

	for _, testCase := range testCases {
		machine := InitVendingMachine(options, price)

		machine.Deposit(testCase.deposit)

		deliveredDrink := machine.Deliver(testCase.choice)
		if deliveredDrink != testCase.drink {
			t.Errorf(
				"%v: TestPricedVendingMachine(%v) == %v, want %v",
				testCase.name,
				testCase.choice,
				deliveredDrink,
				testCase.drink,
			)
		}
	}
}
