package main

type Choice int

const (
	Cola Choice = 0
)

type Drink int

const (
	None Drink = 0
	Coke Drink = 1
)

func InitVendingMachine(choices []Choice) *VendingMachine {
	return &VendingMachine{choices}
}

type VendingMachine struct {
	choices []Choice
}

func (vm *VendingMachine) Deliver(choice Choice) Drink {
	if len(vm.choices) == 0 {
		return None
	}

	return Coke
}