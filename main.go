package main

import (
	"net/http"
	"strconv"

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
