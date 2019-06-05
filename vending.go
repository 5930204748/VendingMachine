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

func (v *VendingMachine) buyDrink(drink string) string {
	m := make(map[string]int)
	m["SD"] = 18
	m["CC"] = 12
	m["DW"] = 7 
	
	price, ok := m[drink]

	if ok {
		if v.totalCoins == price {
			return drink
		}
	}
	return "Add more money"

}

