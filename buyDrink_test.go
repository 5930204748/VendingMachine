package vending

import "testing"

func TestBuy_SD_WithTotalBalance_18_BathShouldReturn_SD(t *testing.T) {
	v := NewVendingMachine()
	expectedResult := "SD"
	v.totalCoins = 18
	
	actualResult := buyDrink("SD")

	if actualResult != expectedResult {
		t.Errorf("%v but got %v", expectedResult, actualResult)
	}
}

func TestBuy_CC_WithTotalBalance_12_BathShouldReturn_CC(t *testing.T) {
	expectedResult := "CC"

	actualResult := buyDrink("CC")

	if actualResult != expectedResult {
		t.Errorf("%v but got %v", expectedResult, actualResult)
	}
}
