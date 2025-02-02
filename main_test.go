package main

import (
	"testing"
)

/*
   There is no need to test for invalid input
   in the receipts as these regular expressions and
   date/time input are validated as the receipts are posted
*/

// TestCountAlphanumericAllAlphanumeric
// Input "Target" Expected Output 6
func TestCountAlphanumericAllAlphanumeric(t *testing.T) {
	retailer := "Target"
	var expected int64 = 6
	var actual int64 = getCountAlphanumericPoints(retailer)
	if actual != expected {
		t.Fatalf(`countAlphanumeric('Target') = %d, expected %d`, actual, expected)
	}
}

// TestCountAlphanumericSemiAlphanumeric
// Input "M&M Corner Market" Expected Output 14
func TestCountAlphanumericSemiAlphanumeric(t *testing.T) {
	retailer := "M&M Corner Market"
	var expected int64 = 14
	var actual int64 = getCountAlphanumericPoints(retailer)
	if actual != expected {
		t.Fatalf(`countAlphanumeric('Target') = %d, expected %d`, actual, expected)
	}
}

// TestCountAlphanumericNoAlphanumeric
// Input "%!@% &&#{} ~+" Expected Output 0
func TestCountAlphanumericNoAlphanumeric(t *testing.T) {
	retailer := "%!@% &&#{} ~+"
	var expected int64 = 0
	var actual int64 = getCountAlphanumericPoints(retailer)
	if actual != expected {
		t.Fatalf(`countAlphanumeric('Target') = %d, expected %d`, actual, expected)
	}
}

// TestRoundDollarIsRoundAmount
// Input "9.00" Expected Output 50
func TestRoundDollarIsRoundAmount(t *testing.T) {
	total := "9.00"
	var expected int64 = 50
	var actual int64 = getRoundDollarPoints(total)
	if actual != expected {
		t.Fatalf(`roundDollar("9.00") = %d, expected %d`, actual, expected)
	}
}

// TestRoundDollarIsNotRoundAmount
// Input "0.57" Expected Output 0
func TestRoundDollarIsNotRoundAmount(t *testing.T) {
	total := "0.57"
	var expected int64 = 0
	var actual int64 = getRoundDollarPoints(total)
	if actual != expected {
		t.Fatalf(`roundDollar("0.57") = %d, expected %d`, actual, expected)
	}
}

// TestMultipleOfQuarterIsNotMultiple
// Input "25.57" Expected Output 0
func TestMultipleOfQuarterIsNotMultiple(t *testing.T) {
	total := "25.57"
	var expected int64 = 0
	var actual int64 = getMultipleOfQuarterPoints(total)
	if actual != expected {
		t.Fatalf(`roundDollar("25.57") = %d, expected %d`, actual, expected)
	}
}

// TestMultipleOfQuarterZeroCent
// Input "10.00" Expected Output 25
func TestMultipleOfQuarterZeroCent(t *testing.T) {
	total := "10.00"
	var expected int64 = 25
	var actual int64 = getMultipleOfQuarterPoints(total)
	if actual != expected {
		t.Fatalf(`roundDollar("10.00") = %d, expected %d`, actual, expected)
	}
}

// TestMultipleOfQuarterTwentyFiveCents
// Input "152.25" Expected Output 25
func TestMultipleOfQuarterTwentyFiveCents(t *testing.T) {
	total := "152.25"
	var expected int64 = 25
	var actual int64 = getMultipleOfQuarterPoints(total)
	if actual != expected {
		t.Fatalf(`roundDollar("152.25") = %d, expected %d`, actual, expected)
	}
}

// TestMultipleOfQuarterFiftyCents
// Input "12.50" Expected Output 25
func TestMultipleOfQuarterFiftyCents(t *testing.T) {
	total := "12.50"
	var expected int64 = 25
	var actual int64 = getMultipleOfQuarterPoints(total)
	if actual != expected {
		t.Fatalf(`roundDollar("12.50") = %d, expected %d`, actual, expected)
	}
}

// TestMultipleOfQuarterSeventyFiveCent
// Input "19.75" Expected Output 25
func TestMultipleOfQuarterSeventyFiveCent(t *testing.T) {
	total := "19.75"
	var expected int64 = 25
	var actual int64 = getMultipleOfQuarterPoints(total)
	if actual != expected {
		t.Fatalf(`roundDollar("19.75") = %d, expected %d`, actual, expected)
	}
}

// TestPairsOneItemLeftover
// Input 5 Pairs Expected Output 10
func TestPairsOneItemLeftover(t *testing.T) {
	var items []Item = []Item{
		{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
		{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
		{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
		{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
		{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"},
	}
	var expected int64 = 10
	actual := getPairsPoints(items)
	if actual != expected {
		t.Fatalf(`getPairsPoints returned %d, expected %d`, actual, expected)
	}
}

// TestPairsAllItemsPaired
// Input 6 Pairs Expected Output 15
func TestPairsAllItemsPaired(t *testing.T) {
	var items []Item = []Item{
		{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
		{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
		{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
		{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
		{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"},
		{ShortDescription: "Lays Potato Chips", Price: "3.99"},
	}
	var expected int64 = 15
	actual := getPairsPoints(items)
	if actual != expected {
		t.Fatalf(`getPairsPoints returned %d, expected %d`, actual, expected)
	}
}

// TestPairsNoPairs
// Input No Pairs Expected Output 0
func TestPairsNoPairs(t *testing.T) {
	var items []Item = []Item{
		{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
	}
	var expected int64 = 0
	actual := getPairsPoints(items)
	if actual != expected {
		t.Fatalf(`getPairsPoints returned %d, expected %d`, actual, expected)
	}
}

// TestTrimmedLengthNoMultiples
// Input No Multiples Expected Output 0
func TestTrimmedLengthNoMultiples(t *testing.T) {
	var items []Item = []Item{
		{ShortDescription: "Mountain Dew 12PK ", Price: "6.49"},
		{ShortDescription: "Emilys Cheese Pizza", Price: "12.25"},
		{ShortDescription: "   Knorr Creamy Chicken", Price: "1.26"},
	}
	var expected int64 = 0
	actual := getItemTrimmedLengthPoints(items)
	if actual != expected {
		t.Fatalf(`getPairsPoints returned %d, expected %d`, actual, expected)
	}
}

// TestTrimmedLength2Multiples
// Input Two Multiples Expected Output 4
func TestTrimmedLength2Multiples(t *testing.T) {
	var items []Item = []Item{
		{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
		{ShortDescription: "   Emils Cheese Pizza   ", Price: "12.25"},
		{ShortDescription: "Knorry Creamy Chicken", Price: "1.26"},
	}
	var expected int64 = 4
	actual := getItemTrimmedLengthPoints(items)
	if actual != expected {
		t.Fatalf(`getPairsPoints returned %d, expected %d`, actual, expected)
	}
}

// TestTrimmedLength1MultipleExtraSpaces
// Input One Multiple Expected Output 6
func TestTrimmedLength1MultipleExtraSpaces(t *testing.T) {
	var items []Item = []Item{
		{ShortDescription: "   Emils  Cheese   Pizza   ", Price: "25.95"},
		{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
		{ShortDescription: "Creamy Chicken", Price: "1.26"},
	}
	var expected int64 = 6
	actual := getItemTrimmedLengthPoints(items)
	if actual != expected {
		t.Fatalf(`getPairsPoints returned %d, expected %d`, actual, expected)
	}
}

// TestPurchaseDateIsEven
// Input "2025-11-30" Expected Output 0
func TestPurchaseDateIsEven(t *testing.T) {
	var purchaseDate = "2025-11-30"
	var expected int64 = 0
	actual := getPurchaseDatePoints(purchaseDate)
	if actual != expected {
		t.Fatalf(`getPurchaseDatePoints("") returned %d, expected %d`, actual, expected)
	}
}

// TestPurchaseDateIsOdd
// Input "2021-09-15" Expected Output 6
func TestPurchaseDateIsOdd(t *testing.T) {
	var purchaseDate = "2021-09-15"
	var expected int64 = 6
	actual := getPurchaseDatePoints(purchaseDate)
	if actual != expected {
		t.Fatalf(`getPurchaseDatePoints("") returned %d, expected %d`, actual, expected)
	}
}

// TestPurchaseDateIsEven
// Input "15:24" Expected Output 10
func TestPurchaseTimeIsBetweenHours(t *testing.T) {
	var purchaseTime = "15:24"
	var expected int64 = 10
	actual := getPurchaseTimePoints(purchaseTime)
	if actual != expected {
		t.Fatalf(`getPurchaseTimePoints("") returned %d, expected %d`, actual, expected)
	}
}

// TestPurchaseTimeIsBeforeHours
// Input "02:24" Expected Output 10
func TestPurchaseTimeIsBeforeHours(t *testing.T) {
	var purchaseTime = "02:24"
	var expected int64 = 0
	actual := getPurchaseTimePoints(purchaseTime)
	if actual != expected {
		t.Fatalf(`getPurchaseTimePoints("") returned %d, expected %d`, actual, expected)
	}
}

// TestPurchaseTimeIsAfterHours
// Input "22:24" Expected Output 0
func TestPurchaseTimeIsAfterHours(t *testing.T) {
	var purchaseTime = "22:24"
	var expected int64 = 0
	actual := getPurchaseTimePoints(purchaseTime)
	if actual != expected {
		t.Fatalf(`getPurchaseTimePoints("") returned %d, expected %d`, actual, expected)
	}
}

// TestPurchaseTimeIsTwoPm
// Input "14:00" Expected Output 10
func TestPurchaseTimeIsTwoPm(t *testing.T) {
	// It's not specified whether or not to include 2pm and 4pm in the points generating time
	// For this project, I'm assuming the 2pm and 4pm are included
	var purchaseTime = "14:00"
	var expected int64 = 10
	actual := getPurchaseTimePoints(purchaseTime)
	if actual != expected {
		t.Fatalf(`getPurchaseTimePoints("") returned %d, expected %d`, actual, expected)
	}
}

// TestPurchaseTimeIsFourPm
// Input "16:00" Expected Output 10
func TestPurchaseTimeIsFourPm(t *testing.T) {
	// It's not specified whether or not to include 2pm and 4pm in the points generating time
	// For this project, I'm assuming the 2pm and 4pm are included
	var purchaseTime = "16:00"
	var expected int64 = 10
	actual := getPurchaseTimePoints(purchaseTime)
	if actual != expected {
		t.Fatalf(`getPurchaseTimePoints("") returned %d, expected %d`, actual, expected)
	}
}
