package main

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

type IdReturn struct {
	ID string `json:"id"`
}

type PointsReturn struct {
	Points int `json:"points"`
}

var receipts map[string]int
var nonAlphaNumericRegex = regexp.MustCompile("[^A-Za-z0-9]+")

func main() {

	// General TODO:
	// 	- Split the points calculation from the validation?
	receipts = make(map[string]int)
	router := mux.NewRouter()

	router.HandleFunc("/receipts/process", func(w http.ResponseWriter, r *http.Request) {

		var receipt Receipt
		err := json.NewDecoder(r.Body).Decode(&receipt)

		// Return an error if we couldn't parse the json
		if err != nil {
			http.Error(w, "The receipt is invalid", http.StatusBadRequest)
			return
		}

		// Validate that all of the data on the receipt is valid before calculating points
		err = validateReceipt(receipt)
		if err != nil {
			http.Error(w, "The receipt is invalid", http.StatusBadRequest)
			return
		}

		// Process the receipt
		receiptID := processReceipt(receipt)

		// Return the receipt ID
		toReturn := IdReturn{
			ID: receiptID,
		}

		json.NewEncoder(w).Encode(toReturn)
	})

	router.HandleFunc("/receipts/{id}/points", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		points, exists := receipts[id]

		// If the specified ID doesn't exist, return an error
		if !exists {
			http.Error(w, "No receipt found for that id", http.StatusBadRequest)
			return
		}

		toReturn := PointsReturn{
			Points: points,
		}

		json.NewEncoder(w).Encode(toReturn)

	})

	http.ListenAndServe(":8080", router)
}

// processReceipt takes a receipt and attempts to calculate the points earned for it.
//
// If successful, it will generate an ID for the receipt and store that ID and the points value into a map.
// If there were errors processing, we will return an error and nothing will be stored.
func processReceipt(receipt Receipt) string {
	// Calculate the points
	points := calculatePoints(receipt)

	// Generate an ID
	id := genID()

	// Store into local memory
	receipts[id] = points

	return id
}

// Returns a new UUID to be used to uniquely identify a receipt
func genID() string {
	id := uuid.NewString() // It's possible for this to panic, I should handle that

	return id
}
