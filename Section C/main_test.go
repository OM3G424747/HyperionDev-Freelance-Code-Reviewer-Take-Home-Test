package main

import (
	"testing"
)

// known valid ISBN-13 example
var validISBN13 string = "9780316066525"

// known invalid ISBN-10 example
var invalidISBN10 string = "0330301824"

// known valid ISBN-10 example
var validISBN10 string = "0316066524"

func TestCheckType(t *testing.T) {
	// test function for checkType function
	valid13 := checkType(validISBN13)
	valid10 := checkType(validISBN10)
	invalidISBN10 := checkType("abc")

	expectedValid13 := 13
	expectedValid10 := 10
	expectedInvalidISBN := -1

	if valid13 != expectedValid13 {
		t.Errorf("Expected Int(%v) is not same as"+
			" actual int (%v)", expectedValid13, valid13)
	}

	if valid10 != expectedValid10 {
		t.Errorf("Expected Int(%v) is not same as"+
			" actual int (%v)", expectedValid10, valid10)
	}

	if invalidISBN10 != expectedInvalidISBN {
		t.Errorf("Expected Int(%v) is not same as"+
			" actual int (%v)", expectedInvalidISBN, invalidISBN10)
	}

}

func TestGetCheckNum(t *testing.T) {
	// test function for getCheckNum function
	valid13 := getCheckNum("9780316066525")
	valid10 := getCheckNum("080442957X")

	expectedValid13 := 5
	expectedValid10 := 10

	if valid13 != expectedValid13 {
		t.Errorf("Expected Int(%v) is not same as"+
			" actual int (%v)", expectedValid13, valid13)
	}

	if valid10 != expectedValid10 {
		t.Errorf("Expected Int(%v) is not same as"+
			" actual int (%v)", expectedValid10, valid10)
	}

}

func TestCheck10(t *testing.T) {
	// test function for check10 function
	valid10 := check10(validISBN10)
	invalid10 := check10(invalidISBN10)

	expectedValid10 := true
	expectedInvalid10 := false

	if valid10 != expectedValid10 {
		t.Errorf("Expected Bool(%v) is not same as"+
			" actual bool (%v)", expectedValid10, valid10)
	}

	if invalid10 != expectedInvalid10 {
		t.Errorf("Expected Bool(%v) is not same as"+
			" actual bool (%v)", expectedInvalid10, invalid10)
	}

}

func TestCheck13(t *testing.T) {
	// test function for check13 function
	valid13 := check13(validISBN13)
	invalid13 := check13("0330301824999")

	expectedValid13 := true
	expectedInvalid13 := false

	if valid13 != expectedValid13 {
		t.Errorf("Expected Bool(%v) is not same as"+
			" actual bool (%v)", expectedValid13, valid13)
	}

	if invalid13 != expectedInvalid13 {
		t.Errorf("Expected Bool(%v) is not same as"+
			" actual bool (%v)", expectedInvalid13, invalid13)
	}

}

func TestGet13Sum(t *testing.T) {
	// test function for get13Sum function
	valid13 := get13Sum(validISBN13)

	expectedValid13 := 90

	if valid13 != expectedValid13 {
		t.Errorf("Expected Int(%v) is not same as"+
			" actual int (%v)", expectedValid13, valid13)
	}

}

func TestToISBN13(t *testing.T) {
	// test function for toISBN13 function
	valid10 := toISBN13(validISBN10)
	altValid10 := toISBN13("080442957X")

	expectedValid10 := "9780316066525"
	expectedAltValid10 := "9780804429573"

	if valid10 != expectedValid10 {
		t.Errorf("Expected String(%v) is not same as"+
			" actual string (%v)", expectedValid10, valid10)
	}

	if altValid10 != expectedAltValid10 {
		t.Errorf("Expected String(%v) is not same as"+
			" actual string (%v)", expectedAltValid10, altValid10)
	}

}
