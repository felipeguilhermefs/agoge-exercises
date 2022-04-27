package main

type Choice int

const (
	Cola        Choice = 0
	FizzyOrange Choice = 1
)

type Drink int

const (
	None  Drink = 0
	Coke  Drink = 1
	Fanta Drink = 2
)

type VendingMachine interface {
	Deliver(choice Choice) Drink
	Deposit(amount int)
}

func InitVendingMachine(options map[Choice]Drink, prices map[Choice]int) VendingMachine {
	return &vendingMachine{options, prices, 0}
}

type vendingMachine struct {
	options map[Choice]Drink
	prices  map[Choice]int
	credits int
}

func (vm *vendingMachine) Deliver(choice Choice) Drink {
	if vm.prices[choice] <= vm.credits {
		return vm.options[choice]
	}
	return None
}

func (vm *vendingMachine) Deposit(amount int) {
	vm.credits = amount
}
