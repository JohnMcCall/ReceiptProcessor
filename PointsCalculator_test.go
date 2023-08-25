package main

import (
	"testing"
)

func checkValue(t *testing.T, caseNum int, expected any, actual any) {
	if expected != actual {
		t.Fatalf(`Failed Case %d. Expected: %v. Actual: %v`, caseNum, expected, actual)
	}
}

func TestGetRetailerPoints(t *testing.T) {
	var expected int
	var actual int
	var testReceipt Receipt

	// Case 1 - Simple Name
	testReceipt.Retailer = "Target"
	expected = 6
	actual = getRetailerPoints(testReceipt)

	checkValue(t, 1, expected, actual)

	// Case 2 - Complex Name
	testReceipt.Retailer = "Bob's Surf 'n Turf"
	expected = 13
	actual = getRetailerPoints(testReceipt)

	checkValue(t, 2, expected, actual)

	// Case 3 - Oops All Nonalphanumeric
	testReceipt.Retailer = "^** '!"
	expected = 0
	actual = getRetailerPoints(testReceipt)

	checkValue(t, 3, expected, actual)

}

func TestGetPricePoints(t *testing.T) {
	var expectedValue int
	var actualValue int
	var testReceipt Receipt

	// Case 1 - Round Value
	testReceipt.Total = "114.00"
	expectedValue = 75
	actualValue = getTotalPoints(testReceipt)

	checkValue(t, 1, expectedValue, actualValue)

	// Case 2 - Change of 0.25
	testReceipt.Total = "35.25"
	expectedValue = 25
	actualValue = getTotalPoints(testReceipt)

	checkValue(t, 2, expectedValue, actualValue)

	// Case 3 - Change of 0.50
	testReceipt.Total = "35.50"
	expectedValue = 25
	actualValue = getTotalPoints(testReceipt)

	checkValue(t, 3, expectedValue, actualValue)

	// Case 4 - Change of 0.75
	testReceipt.Total = "35.75"
	expectedValue = 25
	actualValue = getTotalPoints(testReceipt)

	checkValue(t, 4, expectedValue, actualValue)

	// Case 5 - No points
	testReceipt.Total = "53.99"
	expectedValue = 0
	actualValue = getTotalPoints(testReceipt)

	checkValue(t, 5, expectedValue, actualValue)

}

func TestGetItemPoints(t *testing.T) {
	var expectedValue int
	var actualValue int
	var testReceipt Receipt

	bread := Item{
		Price:            "9.99",
		ShortDescription: "Bread",
	}

	milk := Item{
		Price:            "4.99",
		ShortDescription: "Milk",
	}

	eggs := Item{
		Price:            "5.99",
		ShortDescription: "  Some eggs ",
	}

	bacon := Item{
		Price:            "17.50",
		ShortDescription: "Crispy Bacon",
	}

	// Case 1 - One Item - No Description Points
	testReceipt.Items = append(testReceipt.Items, bread)
	expectedValue = 0
	actualValue = getItemPoints(testReceipt)

	checkValue(t, 1, expectedValue, actualValue)

	// Case 2 - Two Items - No Description Points
	testReceipt.Items = append(testReceipt.Items, milk)
	expectedValue = 5
	actualValue = getItemPoints(testReceipt)

	checkValue(t, 2, expectedValue, actualValue)

	// Case 3 - Three Items - One Items Has Description Points
	testReceipt.Items = append(testReceipt.Items, eggs)
	expectedValue = 7
	actualValue = getItemPoints(testReceipt)

	checkValue(t, 3, expectedValue, actualValue)

	// Case 4 - Four Items - Two Items Have Description Points
	testReceipt.Items = append(testReceipt.Items, bacon)
	expectedValue = 16
	actualValue = getItemPoints(testReceipt)

	checkValue(t, 4, expectedValue, actualValue)

}

func TestGetDatePoints(t *testing.T) {
	var expectedValue int
	var actualValue int
	var testReceipt Receipt

	// Case 1 - Day is even
	testReceipt.PurchaseDate = "2023-08-20"
	expectedValue = 0
	actualValue = getDatePoints(testReceipt)

	checkValue(t, 1, expectedValue, actualValue)

	// Case 2 - Day is odd
	testReceipt.PurchaseDate = "2023-08-15"
	expectedValue = 6
	actualValue = getDatePoints(testReceipt)

	checkValue(t, 2, expectedValue, actualValue)

}

func TestGetTimePoints(t *testing.T) {
	var expectedValue int
	var actualValue int
	var testReceipt Receipt

	// Case 1 - Time is before 2:00 PM
	testReceipt.PurchaseTime = "12:30"
	expectedValue = 0
	actualValue = getTimePoints(testReceipt)

	checkValue(t, 1, expectedValue, actualValue)

	// Case 2 - Time is between 2:00 PM and 4:00 PM
	testReceipt.PurchaseTime = "15:42"
	expectedValue = 10
	actualValue = getTimePoints(testReceipt)

	checkValue(t, 2, expectedValue, actualValue)

	// Case 3 - Time is after 4:00 PM
	testReceipt.PurchaseTime = "20:14"
	expectedValue = 0
	actualValue = getTimePoints(testReceipt)

	checkValue(t, 3, expectedValue, actualValue)

}

func TestIsOdd(t *testing.T) {
	var testNum int
	var expected bool
	var actual bool

	// Case 1 - Number is Odd
	testNum = 5
	expected = true
	actual = isOdd(testNum)

	checkValue(t, 1, expected, actual)

	// Case 2 - Number is Even
	testNum = 18
	expected = false
	actual = isOdd(testNum)

	checkValue(t, 2, expected, actual)

	// Case 3 - Number is Zero
	testNum = 0
	expected = false
	actual = isOdd(testNum)

	checkValue(t, 3, expected, actual)
}
