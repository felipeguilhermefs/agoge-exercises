package main

import "testing"

func TestChoicelessMachineReturnsNothing(t *testing.T) {
	noChoice := map[Choice]Drink{}
	noPrice := map[Choice]int{}
	machine := InitVendingMachine(noChoice, noPrice)

	assetEquals(t, None, machine.Deliver(Cola))
}

func TestDeliversCokeWhenColaIsChosen(t *testing.T) {
	machine := setupFreeVendingMachine()

	assetEquals(t, Coke, machine.Deliver(Cola))
}

func TestDeliversFantaWhenFizzyOrangeIsChosen(t *testing.T) {
	machine := setupFreeVendingMachine()

	assetEquals(t, Fanta, machine.Deliver(FizzyOrange))
}

func TestDeliversNothingWithoutMoney(t *testing.T) {
	machine := setupPricedVendingMachine()

	machine.Deposit(0)

	assetEquals(t, None, machine.Deliver(FizzyOrange))
}

func TestDeliversCokeWhenExactAmountWasDeposited(t *testing.T) {
	machine := setupPricedVendingMachine()

	machine.Deposit(200)

	assetEquals(t, Coke, machine.Deliver(Cola))
}

func TestDeliversFantaWhenDepositIsMoreThanEnough(t *testing.T) {
	machine := setupPricedVendingMachine()

	machine.Deposit(200)

	assetEquals(t, Fanta, machine.Deliver(FizzyOrange))
}

func TestDeliversNothingWhenDepositIsNotEnough(t *testing.T) {
	machine := setupPricedVendingMachine()

	machine.Deposit(100)

	assetEquals(t, None, machine.Deliver(FizzyOrange))
}

func TestDeliversFantaWhenMoneyIsEnoughForFantaButNotForCoke(t *testing.T) {
	machine := setupPricedVendingMachine()

	machine.Deposit(150)

	assetEquals(t, None, machine.Deliver(Cola))
	assetEquals(t, Fanta, machine.Deliver(FizzyOrange))
}

func setupPricedVendingMachine() VendingMachine {
	options := map[Choice]Drink{
		Cola:        Coke,
		FizzyOrange: Fanta,
	}
	prices := map[Choice]int{
		Cola:        200,
		FizzyOrange: 150,
	}
	return InitVendingMachine(options, prices)
}

func setupFreeVendingMachine() VendingMachine {
	options := map[Choice]Drink{
		Cola:        Coke,
		FizzyOrange: Fanta,
	}
	noPrice := map[Choice]int{}
	return InitVendingMachine(options, noPrice)
}

func assetEquals(t *testing.T, expected Drink, received Drink) {
	if expected != received {
		t.Errorf("Expected %v, received %v", expected, received)
	}
}
