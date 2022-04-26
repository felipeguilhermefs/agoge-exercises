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

func InitVendingMachine(options map[Choice]Drink, price int) *VendingMachine {
	return &VendingMachine{options, price}
}

type VendingMachine struct {
	options map[Choice]Drink
	price int
}

func (vm *VendingMachine) Deliver(choice Choice) Drink {
	if vm.price > 0 {
		return None
	}
	return vm.options[choice]
}

func (vm *VendingMachine) Deposit(amount int) {
}
