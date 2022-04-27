package main

import "testing"

func TestChoicelessMachineReturnsNothing(t *testing.T) {
	noChoice := map[Choice]Drink{}
	machine := InitVendingMachine(noChoice, 0)

	assetEquals(t, None, machine.Deliver(Cola))
}

func TestDeliversCokeWhenColaIsChosen(t *testing.T) {
	options := map[Choice]Drink{
		Cola:        Coke,
		FizzyOrange: Fanta,
	}
	machine := InitVendingMachine(options, 0)

	assetEquals(t, Coke, machine.Deliver(Cola))
}

func TestDeliversFantaWhenFizzyOrangeIsChosen(t *testing.T) {
	options := map[Choice]Drink{
		Cola:        Coke,
		FizzyOrange: Fanta,
	}
	machine := InitVendingMachine(options, 0)

	assetEquals(t, Fanta, machine.Deliver(FizzyOrange))
}

func TestDeliversNothingWithoutMoney(t *testing.T) {
	options := map[Choice]Drink{
		Cola:        Coke,
		FizzyOrange: Fanta,
	}
	price := 200
	machine := InitVendingMachine(options, price)

	machine.Deposit(0)

	assetEquals(t, None, machine.Deliver(FizzyOrange))
}

func TestDeliversCokeWhenExactAmountWasDeposited(t *testing.T) {
	options := map[Choice]Drink{
		Cola:        Coke,
		FizzyOrange: Fanta,
	}
	price := 200
	machine := InitVendingMachine(options, price)

	machine.Deposit(200)

	assetEquals(t, Coke, machine.Deliver(Cola))
}

func TestDeliversFantaWhenDepositIsMoreThanEnough(t *testing.T) {
	options := map[Choice]Drink{
		Cola:        Coke,
		FizzyOrange: Fanta,
	}
	price := 150
	machine := InitVendingMachine(options, price)

	machine.Deposit(200)

	assetEquals(t, Fanta, machine.Deliver(FizzyOrange))
}

func TestDeliversNothingWhenDepositIsNotEnough(t *testing.T) {
	options := map[Choice]Drink{
		Cola:        Coke,
		FizzyOrange: Fanta,
	}
	price := 150
	machine := InitVendingMachine(options, price)

	machine.Deposit(100)

	assetEquals(t, None, machine.Deliver(FizzyOrange))
}

func TestDeliversFantaWhenMoneyIsEnoughForFantaButNotForCoke(t *testing.T) {
	options := map[Choice]Drink{
		Cola:        Coke,
		FizzyOrange: Fanta,
	}
	prices := map[Choice]int {
		Cola: 200,
		FizzyOrange: 150,
	}
	machine := InitVendingMachine(options, prices)

	machine.Deposit(150)

	assetEquals(t, None, machine.Deliver(Cola))
	assetEquals(t, Fanta, machine.Deliver(FizzyOrange))
}

func assetEquals(t *testing.T, expected Drink, received Drink) {
	if expected != received {
		t.Errorf("Expected %v, received %v", expected, received)
	}
}
