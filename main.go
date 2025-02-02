package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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

type ReceiptCreatedResponse struct {
	ID string `json:"id"`
}

// Stores the receipts in memory
var receiptsMap map[string]Receipt = make(map[string]Receipt)

// Used to create unique string for receiptId
var receiptNum int = 0

func main() {
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