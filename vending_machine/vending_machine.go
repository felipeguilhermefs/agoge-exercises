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

func InitVendingMachine(options map[Choice]Drink) *VendingMachine {
	return &VendingMachine{options}
}

type VendingMachine struct {
	options map[Choice]Drink
}

func (vm *VendingMachine) Deliver(choice Choice) Drink {
	return vm.options[choice]
}
