package main

import (
	"math"
	"strconv"
	"strings"
	"time"
)

// calculatePoints takes a receipt and calculates the points value of it based on its contents.
//
// Returns the earned points, or an error if we had problems processing
func calculatePoints(receipt Receipt) int {

	retailerPoints := getRetailerPoints(receipt)

	totalPoints := getTotalPoints(receipt)

	itemPoints := getItemPoints(receipt)

	datePoints := getDatePoints(receipt)

	timePoints := getTimePoints(receipt)

	return (retailerPoints + totalPoints + itemPoints + datePoints + timePoints)
}

// getRetailerPoints takes a receipt and calculates the points earned based on the receipt.Retailer.
// One point is awarded for every alphanumeric character in the retailer name.
//
// Returns the earned points.
func getRetailerPoints(receipt Receipt) int {
	strippedRetailerName := nonAlphaNumericRegex.ReplaceAllString(receipt.Retailer, "")
	return len(strippedRetailerName)
}

// getTotalPoints takes a receipt and calculates the points earned based on the receipt.Total.
// 50 points are earned if the total is a round dollar amount with no cents. Additionally,
// 25 points are earned if the total is a multiple of 0.25.
//
// If the total is valid, we will return the earned points. Otherwise, we will return an error.
func getTotalPoints(receipt Receipt) int {
	points := 0
	total, _ := strconv.ParseFloat(receipt.Total, 64)

	// 50 Points if the total is a round dollar amount
	if total == math.Floor(total) {
		points += 50
	}

	// 25 Points if the total is a multiple of 0.25
	if math.Mod(total, 0.25) == 0.0 {
		points += 25
	}

	return points
}

// getItemPoints takes a receipt and calculates the points earned based on the receipt.Items.
// 5 points are earned for every two items on the receipt. Additional points are earned if the
// trimmed length of an item's description is a multiple of 3. In that case the price of that item
// is multiplied by 0.2 and rounded up to the nearest integer. The result of that multiplication
// is added to the points total.
//
// If receipt.Items has one or more items, we will return the earned points. Otherwise, we will return an error.
func getItemPoints(receipt Receipt) int {
	points := 0
	numItems := len(receipt.Items)

	// 5 Points for every two items
	points += 5 * (numItems / 2)

	// If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
	var trimDesc string
	var curItem Item
	for i := 0; i < numItems; i++ {
		curItem = receipt.Items[i]
		trimDesc = strings.Trim(curItem.ShortDescription, " ")

		if math.Mod(float64(len(trimDesc)), 3) == 0 {
			price, _ := strconv.ParseFloat(curItem.Price, 64)

			points += int(math.Ceil(price * 0.2))
		}
	}

	return points
}

// getDatePoints takes a receipt and calculates the points earned based on the receipt.PurchaseDate.
// 6 points are earned if the purchase date is an odd numbered day. 0 points are awarded otherwise.
//
// If the purchase date is valid, we will return the earned points. Otherwise, we will return an error.
func getDatePoints(receipt Receipt) int {
	points := 0
	dateLayout := "2006-01-02"

	purchaseDate, _ := time.Parse(dateLayout, receipt.PurchaseDate)

	if isOdd(purchaseDate.Day()) {
		points += 6
	}

	return points
}

// getTimePoints takes a receipt and calculates the points earned based on the receipt.PurchaseTime.
// 10 points are earned if the purchase time is between 2:00pm and 4:00pm. 0 points are awarded otherwise.
//
// If the purchase time is valid, we will return the earned points. Otherwise, we will return an error.
func getTimePoints(receipt Receipt) int {
	points := 0
	timeLayout := "15:04"
	twoPM, _ := time.Parse(timeLayout, "14:00")
	fourPM, _ := time.Parse(timeLayout, "16:00")

	purchaseTime, _ := time.Parse(timeLayout, receipt.PurchaseTime)

	if purchaseTime.After(twoPM) && purchaseTime.Before(fourPM) {
		points += 10
	}

	return points
}

// isOdd returns true if num is odd, false otherwise
func isOdd(num int) bool {
	return math.Mod(float64(num), float64(2)) != 0
}
