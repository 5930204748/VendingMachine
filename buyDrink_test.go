package vending

import "testing"

func TestBuy_SD_WithTotalBalance_18_BathShouldReturn_SD(t *testing.T) {
	v := NewVendingMachine()
	expectedResult := "SD"
	v.totalCoins = 18

	actualResult := v.buyDrink("SD")

	if actualResult != expectedResult {
		t.Errorf("%v but got %v", expectedResult, actualResult)
	}
}

func TestBuy_CC_WithTotalBalance_12_BathShouldReturn_CC(t *testing.T) {
	expectedResult := "CC"
	v := NewVendingMachine()
	v.totalCoins = 12

	actualResult := v.buyDrink("CC")

	if actualResult != expectedResult {
		t.Errorf("%v but got %v", expectedResult, actualResult)
	}
}
