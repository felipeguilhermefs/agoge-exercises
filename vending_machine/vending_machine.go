package main

type Choice int

const (
	Cola Choice = 0
)

func InitVendingMachine(choices []Choice) *VendingMachine {
	return &VendingMachine{}
}

type VendingMachine struct {
}

func (vm *VendingMachine) Deliver(choice Choice) string {
	return ""
}