package main

func InitVendingMachine(choices []string) *VendingMachine {
	return &VendingMachine{}
}

type VendingMachine struct {
}

func (vm *VendingMachine) Deliver(choice string) string {
	return ""
}