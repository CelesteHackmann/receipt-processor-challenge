package main

import (
	"errors"
	"reflect"
	"time"
)

// Valdiator for Time
func validTime(t interface{}, params string) error {
	format := "15:04"
	// The binding check has already made sure this value is a string
	_, err := time.Parse(format, reflect.ValueOf(t).String())
	if err != nil {
		// Time is Invalid
		return errors.New("invalid time")
	}
	return nil
}

// Validator for Date
func validDate(d interface{}, params string) error {
	format := "2006-01-02"
	// The binding check has already made sure this value is a string
	_, err := time.Parse(format, reflect.ValueOf(d).String())
	if err != nil {
		// Date is Invalid
		return errors.New("invalid date")
	}
	return nil
}