package main

import (
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/gin-gonic/gin"
	"gopkg.in/validator.v2"
)

type Receipt struct {
	Retailer     string `json:"retailer" binding:"required" validate:"regexp=^[\\w\\s\\-&]+$"`
	PurchaseDate string `json:"purchaseDate" binding:"required" validate:"validDate"`
	PurchaseTime string `json:"purchaseTime" binding:"required" validate:"validTime"`
	Items        []Item `json:"items" binding:"required,dive" validate:"min=1"`
	Total        string `json:"total" binding:"required" validate:"regexp=^\\d+\\.\\d{2}$"`
}

type Item struct {
	ShortDescription string `json:"shortDescription" binding:"required" validate:"regexp=^[\\w\\s\\-&]+$"`
	Price            string `json:"price" binding:"required" validate:"regexp=^\\d+\\.\\d{2}$"`
}

type ReceiptCreatedResponse struct {
	ID string `json:"id"`
}

type PointsGeneratedResponse struct {
	Points int64 `json:"points"`
}

// Stores the receipts in memory
var receiptsMap map[string]Receipt = make(map[string]Receipt)

// Used to create unique string for receiptId
var receiptNum int = 0

func main() {
	// Add validation functions for Time and Date
	validator.SetValidationFunc("validTime", validTime)
	validator.SetValidationFunc("validDate", validDate)

	// Create Gin router
	router := gin.Default()

	// Define the api paths
	router.POST("/receipts/process", processReceipt)
	router.GET("/receipts/:id/points", getPoints)

	// Start the server
	router.Run("localhost:8080")
}

// processReceipt validate the JSON body, assigns the receipt a unique id, adds the Receipt to the map, and gives the id to the response
func processReceipt(c *gin.Context) {
	var newReceipt Receipt

	// Check if the requestBody and resulting Receipt is valid, if not it returns 400 BadRequest
	if err := c.BindJSON(&newReceipt); err != nil {
		c.String(http.StatusBadRequest, "The receipt is invalid.")
		return
	}
	// Validate the struct
	if err := validator.Validate(newReceipt); err != nil {
		c.String(http.StatusBadRequest, "The receipt is invalid.")
		return
	}

	// Generates a unique id and save receipt
	// This works since the data is not persistant
	receiptNum += 1
	var receiptId string = "Receipt" + strconv.Itoa(receiptNum)
	receiptsMap[receiptId] = newReceipt

	// Add id to the Response
	response := ReceiptCreatedResponse{
		ID: receiptId,
	}
	c.JSON(http.StatusOK, response)
}

// getPoints calculates and returns the amount of points awarded for a receipt given the receiptId
func getPoints(c *gin.Context) {
	// Check if the receiptId is valid, if not return a 404 NotFound
	var receiptId = c.Param("id")
	receipt, ok := receiptsMap[receiptId]
	if !ok {
		c.String(http.StatusNotFound, "No receipt found for that ID.")
		return
	}

	// Calcuate and add points to context response
	var points int64 = calcuatePoints(receipt)
	response := PointsGeneratedResponse{
		Points: points,
	}
	c.JSON(http.StatusOK, response)
}

// CalcualtePoints gets and adds up all the points for the receipt
func calcuatePoints(receipt Receipt) int64 {
	var points int64 = 0
	points += getCountAlphanumericPoints(receipt.Retailer)

	points += getRoundDollarPoints(receipt.Total)

	points += getMultipleOfQuarterPoints(receipt.Total)

	points += getPairsPoints(receipt.Items)

	points += getItemTrimmedLengthPoints(receipt.Items)

	points += getPurchaseDatePoints(receipt.PurchaseDate)

	points += getPurchaseTimePoints(receipt.PurchaseTime)

	return points
}

// One point for every alphanumeric character in the retailer name
func getCountAlphanumericPoints(retailer string) int64 {
	var points int64 = 0
	for _, r := range retailer {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			points += 1
		}
	}
	return points
}

// 50 points if the total is a round dollar amount with no cents
func getRoundDollarPoints(total string) int64 {
	result := strings.SplitAfter(total, ".")
	if result[1] == "00" {
		return 50
	} else {
		return 0
	}
}

// 25 points if the total is a multiple of 0.25
func getMultipleOfQuarterPoints(total string) int64 {
	result := strings.SplitAfter(total, ".")
	centAmount, _ := strconv.Atoi(result[1])
	// If cents is 00, 25, 50, 75 then add the points
	if centAmount%25 == 0 {
		return 25
	} else {
		return 0
	}
}

// 5 points for every two items on the receipt
func getPairsPoints(items []Item) int64 {
	numItems := len(items)
	numPairs := numItems / 2
	return int64(numPairs * 5)
}

// If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned
func getItemTrimmedLengthPoints(items []Item) int64 {
	var points int64 = 0
	// for each item
	for _, item := range items {
		// use strings.TrimSpace to remove the leading and trailing whitespace
		trimmedItemDescription := strings.TrimSpace(item.ShortDescription)
		// Get lgenth
		trimmedLength := len(trimmedItemDescription)
		// If trimmed length is a multiple of 3
		if trimmedLength%3 == 0 {
			// multiple the price by .2
			val, _ := strconv.ParseFloat(item.Price, 64)
			// round up to nearest integer
			// add this number to points
			points += int64(math.Ceil(val * .2))
		}
	}
	return points
}

// 6 points if the day in the purchase date is odd
func getPurchaseDatePoints(purchaseDate string) int64 {
	format := "2006-01-02"
	// Already checked for valid date with validator
	date, _ := time.Parse(format, purchaseDate)
	if date.Day()%2 == 1 {
		return 6
	} else {
		return 0
	}
}

// 10 points if the time of purchase is after 2:00pm and before 4:00pm
func getPurchaseTimePoints(purchaseTime string) int64 {
	format := "15:04"
	// Already checked for valid time with validator
	pTime, _ := time.Parse(format, purchaseTime)
	twoPm, _ := time.Parse(format, "14:00")
	fourPm, _ := time.Parse(format, "16:00")
	if isBetweenTimeRange(pTime, twoPm, fourPm) {
		return 10
	} else {
		return 0
	}
}

// Check to see if given time is between 2pm and 4pm, inclusive of 2pm and 4pm
func isBetweenTimeRange(pTime time.Time, firstTime time.Time, secondTime time.Time) bool {
	if (pTime.After(firstTime) && pTime.Before(secondTime)) || (pTime.Equal(firstTime) || (pTime.Equal(secondTime))) {
		return true
	}
	return false
}