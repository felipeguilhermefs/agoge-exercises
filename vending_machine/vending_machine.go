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

type Product struct {
	drink Drink
	price int
}

type VendingMachine interface {
	Deliver(choice Choice) Drink
	Deposit(amount int)
}

func InitVendingMachine(products map[Choice]Product) VendingMachine {
	return &vendingMachine{products, 0}
}

type vendingMachine struct {
	products map[Choice]Product
	credits  int
}

func (vm *vendingMachine) Deliver(choice Choice) Drink {
	product := vm.products[choice]
	if product.price <= vm.credits {
		return product.drink
	}
	return None
}

func (vm *vendingMachine) Deposit(amount int) {
	vm.credits = amount
}
