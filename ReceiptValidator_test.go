package main

import "testing"

func checkError(t *testing.T, caseNum int, errorExpected bool, err error) {
	//Fail if we have an error we weren't expecting OR don't have an error we were expecting
	if ((err != nil) && !errorExpected) || ((err == nil) && errorExpected) {
		t.Fatalf(`Failed Error Case %d. Error Expected?: %v. Error Value: %v`, caseNum, errorExpected, err)
	}
}

func TestValidateRetailer(t *testing.T) {
	var expected bool
	var actual error
	var testRetailer string

	// Case 1 - Valid Retailer
	testRetailer = "Target"
	expected = false
	actual = validateRetailer(testRetailer)

	checkError(t, 1, expected, actual)

	// Case 2 - Missing Retailer
	testRetailer = ""
	expected = true
	actual = validateRetailer(testRetailer)

	checkError(t, 2, expected, actual)
}

func TestValidateTotal(t *testing.T) {
	var expected bool
	var actual error
	var testTotal string

	// Case 1 - Valid Total
	testTotal = "16.58"
	expected = false
	actual = validateTotal(testTotal)

	checkError(t, 1, expected, actual)

	// Case 2 - Negative Total
	testTotal = "-15.49"
	expected = true
	actual = validateTotal(testTotal)

	checkError(t, 2, expected, actual)

	// Case 3 - Missing Total
	testTotal = ""
	expected = true
	actual = validateTotal(testTotal)

	checkError(t, 3, expected, actual)
}

func TestValidateItems(t *testing.T) {
	var expected bool
	var actual error
	var testItems []Item

	validItem := Item{
		Price:            "9.99",
		ShortDescription: "Bread",
	}

	invalidItem := Item{
		Price:            "4.99",
		ShortDescription: "",
	}

	// Case 1 - Empty List
	expected = true
	actual = validateItems(testItems)

	checkError(t, 1, expected, actual)

	// Case 2 - Valid List
	testItems = append(testItems, validItem)
	expected = false
	actual = validateItems(testItems)

	checkError(t, 2, expected, actual)

	// Case 3 - Invalid List
	testItems = append(testItems, invalidItem)
	expected = true
	actual = validateItems(testItems)

	checkError(t, 3, expected, actual)
}

func TestValidateItem(t *testing.T) {
	var expected bool
	var actual error
	var testItem Item

	// Case 1 - Valid Item
	testItem.Price = "16.58"
	testItem.ShortDescription = "Gum"
	expected = false
	actual = validateItem(testItem)

	checkError(t, 1, expected, actual)

	// Case 2 - Negative Price
	testItem.Price = "-16.58"
	testItem.ShortDescription = "Gum"
	expected = true
	actual = validateItem(testItem)

	checkError(t, 2, expected, actual)

	// Case 3 - Missing Price
	testItem.Price = ""
	testItem.ShortDescription = "Gum"
	expected = true
	actual = validateItem(testItem)

	checkError(t, 3, expected, actual)

	// Case 4 - Missing Description
	testItem.Price = "16.58"
	testItem.ShortDescription = ""
	expected = true
	actual = validateItem(testItem)

	checkError(t, 4, expected, actual)
}

func TestValidatePurchaseDate(t *testing.T) {
	var expected bool
	var actual error
	var testDate string

	// Case 1 - Valid Date
	testDate = "2023-08-20"
	expected = false
	actual = validatePurchaseDate(testDate)

	checkError(t, 1, expected, actual)

	// Case 2 - Invalid Date
	testDate = "2023-15-04"
	expected = true
	actual = validatePurchaseDate(testDate)

	checkError(t, 2, expected, actual)
}

func TestValidatePurchaseTime(t *testing.T) {
	var expected bool
	var actual error
	var testTime string

	// Case 1 - Valid Date
	testTime = "15:22"
	expected = false
	actual = validatePurchaseTime(testTime)

	checkError(t, 1, expected, actual)

	// Case 2 - Invalid Time
	testTime = "4:00 PM"
	expected = true
	actual = validatePurchaseTime(testTime)

	checkError(t, 2, expected, actual)
}
