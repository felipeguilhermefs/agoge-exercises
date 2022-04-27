package main

import "testing"

func TestChoicelessMachineReturnsNothing(t *testing.T) {
	noProducts := map[Choice]Product{}
	machine := InitVendingMachine(noProducts)

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

func TestDeliversCokeAfterDepositIsEnough(t *testing.T) {
	machine := setupPricedVendingMachine()

	machine.Deposit(150)
	assetEquals(t, None, machine.Deliver(Cola))

	machine.Deposit(50)
	assetEquals(t, Coke, machine.Deliver(Cola))
}

func setupPricedVendingMachine() VendingMachine {
	products := map[Choice]Product{
		Cola:        {Coke, 200},
		FizzyOrange: {Fanta, 150},
	}
	return InitVendingMachine(products)
}

func setupFreeVendingMachine() VendingMachine {
	products := map[Choice]Product{
		Cola:        {Coke, 0},
		FizzyOrange: {Fanta, 0},
	}
	return InitVendingMachine(products)
}

func assetEquals(t *testing.T, expected Drink, received Drink) {
	if expected != received {
		t.Errorf("Expected %v, received %v", expected, received)
	}
}
