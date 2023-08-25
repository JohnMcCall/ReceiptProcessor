package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

// Validates all of the required information on the receipt
//
// Returns an error if any part of the receipt is invalid, nil otherwise
func validateReceipt(receipt Receipt) error {
	var err error

	err = validateRetailer(receipt.Retailer)
	if err != nil {
		return err
	}

	err = validateTotal(receipt.Total)
	if err != nil {
		return err
	}

	err = validateItems(receipt.Items)
	if err != nil {
		return err
	}

	err = validatePurchaseDate(receipt.PurchaseDate)
	if err != nil {
		return err
	}

	err = validatePurchaseTime(receipt.PurchaseTime)
	if err != nil {
		return err
	}

	return nil
}

// Validates the retailer on the receipt
//
// Returns an error if the retailer isn't set, nil otherwise
func validateRetailer(retailer string) error {
	// The pattern in the API implies no spaces are allowed, but the example have spaces.
	// I'm going to assume the retailer name is valid as long as it's not null
	var err error

	if retailer == "" {
		err = errors.New("Retailer is invalid")
	}

	return err
}

// Validates the total on the receipt
//
// Returns an error if the total can't be parsed into a float or if it's negative, nil otherwise
func validateTotal(total string) error {

	floatTotal, err := strconv.ParseFloat(total, 64)

	// If the total is less than or equal to zero, return an error.
	if (err != nil) || (floatTotal < 0) {
		err = fmt.Errorf("Total: '%s' is invalid", total)
	}

	return err
}

// Validates the list of purchased items
//
// Returns an error if the list is empty, or if an item in the list isn't valid
func validateItems(items []Item) error {
	var err error

	// Check that we have items
	if len(items) < 1 {
		err = fmt.Errorf("Items are invalid")
		return err
	}

	// Validate each item
	for i := 1; i < len(items); i++ {
		err = validateItem(items[i])
		if err != nil {
			break
		}
	}

	return err
}

// Validates a single item
//
// Returns an error if the price can't be parsed or is negative, or if the description is missing. Returns nil otherwise
func validateItem(item Item) error {
	var err error

	// Validate Price
	itemPrice, err := strconv.ParseFloat(item.Price, 64)
	if (err != nil) || (itemPrice < 0) {
		err = fmt.Errorf("Price: '%s' is invalid", item.Price)
		return err
	}

	// Validate description
	if item.ShortDescription == "" {
		err = fmt.Errorf("Short Description is invalid")
	}

	return err
}

// Validates the purchase date on the receipt
//
// Returns an error if the purchased date isn't a valid date, nil otherise
func validatePurchaseDate(purchaseDate string) error {
	dateLayout := "2006-01-02"
	_, err := time.Parse(dateLayout, purchaseDate)
	return err
}

// Validates the purchase time on the receipt
//
// Returns an error if the purchased time isn't a valid time, nil otherise
func validatePurchaseTime(purchaseTime string) error {
	timeLayout := "15:04"
	_, err := time.Parse(timeLayout, purchaseTime)
	return err
}
