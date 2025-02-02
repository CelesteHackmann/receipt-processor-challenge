package main

import (
	"testing"
	"gopkg.in/validator.v2"
)

// TestValidTimeWithValidTimeandDate
// Valid Time and Date Expect No Error
func TestValidTimeWithValidTimeandDate(t *testing.T) {
	receipt := Receipt{
		Retailer: "Target",
		PurchaseDate: "2024-12-12",
		PurchaseTime: "05:31",
		Items: []Item{
			{ShortDescription: "Mountain Dew 12PK",Price: "6.49"},
			{ShortDescription: "Knorr Creamy Chicken",Price: "1.26"},
		},
		Total: "3.00",
	}

	validator.SetValidationFunc("validTime", validTime)
	validator.SetValidationFunc("validDate", validDate)
	err := validator.Validate(receipt)
	if err != nil {
		t.Fatalf(`validTime("05:31") = %v, expected no error`, err)
	}
}

// TestValidTimeWithValidTimeInvalidDate
// Invalid Date Expect Error
func TestValidTimeWithValidTimeInvalidDate(t *testing.T) {
	receipt := Receipt{
		Retailer: "Target",
		PurchaseDate: "2000-02-30",
		PurchaseTime: "05:31",
		Items: []Item{
			{ShortDescription: "Mountain Dew 12PK",Price: "6.49"},
			{ShortDescription: "Knorr Creamy Chicken",Price: "1.26"},
		},
		Total: "3.00",
	}

	validator.SetValidationFunc("validTime", validTime)
	validator.SetValidationFunc("validDate", validDate)
	err := validator.Validate(receipt)
	if err == nil {
		t.Fatalf(`validDate("2000-02-30") = %v, expected invalid date error`, err)
	}
}

// TestValidTimeWithInvalidTimealidDate
// Invalid Time Expect Error
func TestValidTimeWithInvalidTimeValidDate(t *testing.T) {
	receipt := Receipt{
		Retailer: "Target",
		PurchaseDate: "200-01-30",
		PurchaseTime: "32:00",
		Items: []Item{
			{ShortDescription: "Mountain Dew 12PK",Price: "6.49"},
			{ShortDescription: "Knorr Creamy Chicken",Price: "1.26"},
		},
		Total: "3.00",
	}

	validator.SetValidationFunc("validTime", validTime)
	validator.SetValidationFunc("validDate", validDate)
	err := validator.Validate(receipt)
	if err == nil {
		t.Fatalf(`validDate("200-01-30") = %v, expected invalid date and invalid time error`, err)
	}
}


// TestValidTimeWithInvalidTimealidDate
// Invalid Time and Invalid Date Expect Error
func TestValidTimeWithInvalidTimeInvalidDate(t *testing.T) {
	receipt := Receipt{
		Retailer: "Target",
		PurchaseDate: "200-01-30",
		PurchaseTime: "32:00",
		Items: []Item{
			{ShortDescription: "Mountain Dew 12PK",Price: "6.49"},
			{ShortDescription: "Knorr Creamy Chicken",Price: "1.26"},
		},
		Total: "3.00",
	}

	validator.SetValidationFunc("validTime", validTime)
	validator.SetValidationFunc("validDate", validDate)
	err := validator.Validate(receipt)
	if err == nil {
		t.Fatalf(`validDate("200-01-30") = %v, expected invalid date and invalid time error`, err)
	}
}