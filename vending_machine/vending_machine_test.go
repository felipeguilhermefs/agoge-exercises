package main

import "testing"

func TestChoicelessMachineReturnsNothing(t *testing.T) {
	noChoice := map[Choice]Drink{}
	machine := InitVendingMachine(noChoice, 0)

	deliveredDrink := machine.Deliver(Cola)

	if deliveredDrink != None {
		t.Errorf("(%v) == %v, want %v", Cola, deliveredDrink, None)
	}
}

func TestDeliversCokeWhenColaIsChosen(t *testing.T) {
	options := map[Choice]Drink{
		Cola:        Coke,
		FizzyOrange: Fanta,
	}
	machine := InitVendingMachine(options, 0)

	deliveredDrink := machine.Deliver(Cola)

	if deliveredDrink != Coke {
		t.Errorf("(%v) == %v, want %v", Cola, deliveredDrink, Coke)
	}
}

func TestDeliversFantaWhenFizzyOrangeIsChosen(t *testing.T) {
	options := map[Choice]Drink{
		Cola:        Coke,
		FizzyOrange: Fanta,
	}
	machine := InitVendingMachine(options, 0)

	deliveredDrink := machine.Deliver(FizzyOrange)

	if deliveredDrink != Fanta {
		t.Errorf("(%v) == %v, want %v", Cola, deliveredDrink, Fanta)
	}
}

func TestDeliversNothingWithoutMoney(t *testing.T) {
	options := map[Choice]Drink{
		Cola:        Coke,
		FizzyOrange: Fanta,
	}
	price := 200
	machine := InitVendingMachine(options, price)

	machine.Deposit(0)

	deliveredDrink := machine.Deliver(FizzyOrange)

	if deliveredDrink != None {
		t.Errorf("(%v) == %v, want %v", Cola, deliveredDrink, None)
	}
}

func TestDeliversCokeWhenExactAmountWasDeposited(t *testing.T) {
	options := map[Choice]Drink{
		Cola:        Coke,
		FizzyOrange: Fanta,
	}
	price := 200
	machine := InitVendingMachine(options, price)

	machine.Deposit(200)

	deliveredDrink := machine.Deliver(Cola)

	if deliveredDrink != Coke {
		t.Errorf("(%v) == %v, want %v", Cola, deliveredDrink, Coke)
	}
}

func TestDeliversFantaWhenDepositIsMoreThanEnough(t *testing.T) {
	options := map[Choice]Drink{
		Cola:        Coke,
		FizzyOrange: Fanta,
	}
	price := 150
	machine := InitVendingMachine(options, price)

	machine.Deposit(200)

	deliveredDrink := machine.Deliver(FizzyOrange)

	if deliveredDrink != Fanta {
		t.Errorf("(%v) == %v, want %v", FizzyOrange, deliveredDrink, Fanta)
	}
}

func TestDeliversNothingWhenDepositIsNotEnough(t *testing.T) {
	options := map[Choice]Drink{
		Cola:        Coke,
		FizzyOrange: Fanta,
	}
	price := 150
	machine := InitVendingMachine(options, price)

	machine.Deposit(100)

	deliveredDrink := machine.Deliver(FizzyOrange)

	if deliveredDrink != None {
		t.Errorf("(%v) == %v, want %v", FizzyOrange, deliveredDrink, None)
	}
}
