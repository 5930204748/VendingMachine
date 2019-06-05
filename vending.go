package vending

type VendingMachine struct {
	totalCoins int
	m          map[string]int
}

func (v VendingMachine) showTotalBalance() int {
	return v.totalCoins
}

func (v *VendingMachine) insertCoins(coin string) {
	elem := v.m[coin]
	v.totalCoins += elem

}

func NewVendingMachine() VendingMachine {
	m := make(map[string]int)
	m["T"] = 10
	m["F"] = 5
	m["TW"] = 2
	m["O"] = 1
	return VendingMachine{m: m}
}
