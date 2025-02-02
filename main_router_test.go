package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gopkg.in/validator.v2"
)

// The following tests test the HTTP response from the api requests

// Setup and Teardown functions used at the beginning and end of each test case
func setup() {
	validator.SetValidationFunc("validTime", validTime)
	validator.SetValidationFunc("validDate", validDate)
}

func teardown() {
	for k := range receiptsMap {
		delete(receiptsMap, k)
	}
}

// Receipts
var validReceipt1 Receipt = Receipt{
	Retailer:     "Target",
	PurchaseDate: "2022-01-01",
	PurchaseTime: "02:01",
	Items: []Item{
		{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
		{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
		{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
		{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
		{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"},
	},
	Total: "35.35",
}

var validReceipt2 Receipt = Receipt{
	Retailer:     "M&M Corner Market",
	PurchaseDate: "2022-03-20",
	PurchaseTime: "14:33",
	Total:        "9.00",
	Items: []Item{
		{ShortDescription: "Gatorade", Price: "2.25"},
		{ShortDescription: "Gatorade", Price: "2.25"},
		{ShortDescription: "Gatorade", Price: "2.25"},
		{ShortDescription: "Gatorade", Price: "2.25"},
	},
}

var validReceipt3 Receipt = Receipt{
	Retailer:     "Target-Kroger",
	PurchaseDate: "2022-10-03",
	PurchaseTime: "18:00",
	Items: []Item{
		{ShortDescription: "Pepsi - 12-oz", Price: "1.25"},
		{ShortDescription: "   Powerade Red  ", Price: "5.00"},
		{ShortDescription: "Cheese Pizza", Price: "10.25"},
		{ShortDescription: "Super-Duper  Hot  Cocoa Mix   ", Price: "20.25"},
	},
	Total: "35.75",
}

var receiptInvalidRetailer Receipt = Receipt{
	Retailer:     "Tar.get",
	PurchaseDate: "2022-01-02",
	PurchaseTime: "13:13",
	Items: []Item{
		{ShortDescription: "Pepsi - 12-oz", Price: "1.25"},
	},
	Total: "1.25",
}

var receiptInvalidDate Receipt = Receipt{
	Retailer:     "Target",
	PurchaseDate: "2022-19-02",
	PurchaseTime: "13:13",
	Items: []Item{
		{ShortDescription: "Pepsi - 12-oz", Price: "1.25"},
	},
	Total: "1.25",
}

var receiptInvalidPurchaseTime Receipt = Receipt{
	Retailer:     "Target",
	PurchaseDate: "2002-10-02",
	PurchaseTime: "05:91",
	Items: []Item{
		{ShortDescription: "Pepsi - 12-oz", Price: "1.25"},
	},
	Total: "1.25",
}

var receiptNoItems Receipt = Receipt{
	Retailer:     "Meijer",
	PurchaseDate: "2002-10-02",
	PurchaseTime: "05:51",
	Items:        []Item{},
	Total:        "0.00",
}

var receiptInvalidItemDescription Receipt = Receipt{
	Retailer:     "Meijer",
	PurchaseDate: "2002-10-02",
	PurchaseTime: "05:41",
	Items: []Item{
		{ShortDescription: "Pepsi - 12-oz", Price: "1.25"},
		{ShortDescription: "Not a good, description", Price: "51.05"},
	},
	Total: "52.30",
}

var receiptInvalidItemPrice Receipt = Receipt{
	Retailer:     "Kroger",
	PurchaseDate: "1995-08-02",
	PurchaseTime: "23:41",
	Items: []Item{
		{ShortDescription: "Pepsi - 12-oz", Price: "1.25"},
		{ShortDescription: "A better description", Price: "51.500"},
	},
	Total: "51.50",
}

var receiptInvalidTotal Receipt = Receipt{
	Retailer:     "Kroger",
	PurchaseDate: "1995-08-02",
	PurchaseTime: "23:41",
	Items: []Item{
		{ShortDescription: "Pepsi - 12-oz", Price: "1.25"},
		{ShortDescription: "A better description", Price: "51.50"},
	},
	Total: "51.500",
}

// TESTS

// TestProcessReceiptMultipleValidReceipts
func TestProcessReceiptMultipleValidReceipts(t *testing.T) {
	setup()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &http.Request{
		Header: make(http.Header),
	}
	c.Request.Method = "POST"
	c.Request.Header.Set("Content-Type", "application/json")
	jsonbytes, err := json.Marshal(validReceipt1)
	if err != nil {
		panic(err)
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))

	processReceipt(c)
	assert.EqualValues(t, http.StatusOK, w.Code)
	expectedResponse, _ := json.Marshal(ReceiptCreatedResponse{ID: "Receipt1"})
	actualResponse := w.Body.String()
	assert.EqualValues(t, expectedResponse, actualResponse)

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
	c.Request = &http.Request{
		Header: make(http.Header),
	}
	c.Request.Method = "POST"
	c.Request.Header.Set("Content-Type", "application/json")
	jsonbytes, err = json.Marshal(validReceipt2)
	if err != nil {
		panic(err)
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))

	processReceipt(c)
	assert.EqualValues(t, http.StatusOK, w.Code)
	expectedResponse, _ = json.Marshal(ReceiptCreatedResponse{ID: "Receipt2"})
	actualResponse = w.Body.String()
	assert.EqualValues(t, expectedResponse, actualResponse)

	teardown()
}

// TestProcessReceiptInvalidRetailer
func TestProcessReceiptInvalidRetailer(t *testing.T) {
	setup()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &http.Request{
		Header: make(http.Header),
	}
	c.Request.Method = "POST"
	c.Request.Header.Set("Content-Type", "application/json")
	jsonbytes, err := json.Marshal(receiptInvalidRetailer)
	if err != nil {
		panic(err)
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))

	processReceipt(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "The receipt is invalid.", w.Body.String())

	teardown()
}

// TestProcessReceiptInvalidPurchaseDate
func TestProcessReceiptInvalidPurchaseDate(t *testing.T) {
	setup()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &http.Request{
		Header: make(http.Header),
	}
	c.Request.Method = "POST"
	c.Request.Header.Set("Content-Type", "application/json")
	jsonbytes, err := json.Marshal(receiptInvalidDate)
	if err != nil {
		panic(err)
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))

	processReceipt(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "The receipt is invalid.", w.Body.String())

	teardown()
}

// TestProcessReceiptInvalidPurchaseTime
func TestProcessReceiptInvalidPurchaseTime(t *testing.T) {
	setup()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &http.Request{
		Header: make(http.Header),
	}
	c.Request.Method = "POST"
	c.Request.Header.Set("Content-Type", "application/json")
	jsonbytes, err := json.Marshal(receiptInvalidPurchaseTime)
	if err != nil {
		panic(err)
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))

	processReceipt(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "The receipt is invalid.", w.Body.String())

	teardown()
}

// TestProcessReceiptNoItems
func TestProcessReceiptInvalidNoItems(t *testing.T) {
	setup()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &http.Request{
		Header: make(http.Header),
	}
	c.Request.Method = "POST"
	c.Request.Header.Set("Content-Type", "application/json")
	jsonbytes, err := json.Marshal(receiptNoItems)
	if err != nil {
		panic(err)
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))

	processReceipt(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "The receipt is invalid.", w.Body.String())

	teardown()
}

// TestProcessReceiptInvalidItemDescription
func TestProcessReceiptInvalidItemDescription(t *testing.T) {
	setup()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &http.Request{
		Header: make(http.Header),
	}
	c.Request.Method = "POST"
	c.Request.Header.Set("Content-Type", "application/json")
	jsonbytes, err := json.Marshal(receiptInvalidItemDescription)
	if err != nil {
		panic(err)
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))

	processReceipt(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "The receipt is invalid.", w.Body.String())

	teardown()
}

// TestProcessReceiptInvalidItemPrice
func TestProcessReceiptInvalidItemPrice(t *testing.T) {
	setup()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &http.Request{
		Header: make(http.Header),
	}
	c.Request.Method = "POST"
	c.Request.Header.Set("Content-Type", "application/json")
	jsonbytes, err := json.Marshal(receiptInvalidItemPrice)
	if err != nil {
		panic(err)
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))

	processReceipt(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "The receipt is invalid.", w.Body.String())

	teardown()
}

// TestProcessReceiptInvalidTotal
func TestProcessReceiptInvalidTotal(t *testing.T) {
	setup()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &http.Request{
		Header: make(http.Header),
	}
	c.Request.Method = "POST"
	c.Request.Header.Set("Content-Type", "application/json")
	jsonbytes, err := json.Marshal(receiptInvalidTotal)
	if err != nil {
		panic(err)
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))

	processReceipt(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "The receipt is invalid.", w.Body.String())

	teardown()
}

// TestCalculatePointsInvalidId
func TestCalculatePointsInvalidId(t *testing.T) {
	setup()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &http.Request{
		Header: make(http.Header),
	}
	c.Request.Method = "GET"
	c.Params = []gin.Param{
		{
			Key:   "id",
			Value: "Receipt10",
		},
	}

	getPoints(c)
	expectedString := "No receipt found for that ID."
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, expectedString, w.Body.String())

	teardown()
}

// TestCalculatePointsReceipt1
func TestCalculatePointsReceipt1(t *testing.T) {
	setup()
	expectedResponse, _ := json.Marshal(PointsGeneratedResponse{Points: 28})
	receiptsMap["Receipt1"] = validReceipt1
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &http.Request{
		Header: make(http.Header),
	}
	c.Request.Method = "GET"
	c.Params = []gin.Param{
		{
			Key:   "id",
			Value: "Receipt1",
		},
	}

	getPoints(c)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, string(expectedResponse), w.Body.String())

	teardown()
}

// TestCalculatePointsReceipt2
func TestCalculatePointsReceipt2(t *testing.T) {
	setup()
	expectedResponse, _ := json.Marshal(PointsGeneratedResponse{Points: 109})
	
	receiptsMap["Receipt1"] = validReceipt2
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &http.Request{
		Header: make(http.Header),
	}
	c.Request.Method = "GET"
	c.Params = []gin.Param{
		{
			Key:   "id",
			Value: "Receipt1",
		},
	}

	getPoints(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, string(expectedResponse), w.Body.String())

	teardown()
}

// // TestCalculatePointsReceipt3
func TestCalculatePointsReceipt3(t *testing.T) {
	setup()
	expectedResponse, _ := json.Marshal(PointsGeneratedResponse{Points: 62})

	receiptsMap["Receipt12"] = validReceipt3
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &http.Request{
		Header: make(http.Header),
	}
	c.Request.Method = "GET"
	c.Params = []gin.Param{
		{
			Key:   "id",
			Value: "Receipt12",
		},
	}

	getPoints(c)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, string(expectedResponse), w.Body.String())

	teardown()
}