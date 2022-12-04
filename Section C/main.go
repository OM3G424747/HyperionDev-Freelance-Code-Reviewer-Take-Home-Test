package main

import (
	"bufio"
	"fmt"
	"os"
	cast "strconv"
	"strings"
)

func checkType(isbn string) int {
	// checks if the string is ISBN-10 or ISBN-13 based on lenght
	// returns 10 for ISBN-10 or 13 for ISBN-13
	// returns -1 if it's not a valid length
	if len(isbn) == 10 {
		return 10
	}

	if len(isbn) == 13 {
		return 13
	}

	return -1
}

func getCheckNum(isbn string) int {
	// returns int matching check number value
	checkNumber := string(isbn[len(isbn)-1])
	checkNumber = strings.ToLower(checkNumber)

	if checkNumber == "x" {
		checkNumber = "10"
	}

	value, _ := cast.Atoi(checkNumber)
	return value
}

func check10(isbn string) bool {
	// confirms if ISBN-10 string is valid or not
	// returns true for valid ISBN
	// returns false if ISBN is not valid
	var bookSum int = 0
	var value int
	var numberIndex int = 0

	// adds check number to books sum
	bookSum += getCheckNum(isbn)

	// itterates through 10 to 1 to multiply each number and add it to the bookSum
	for i := 10; i >= 2; i-- {
		// casts char byte value to string to get UTF-8 value
		var currentNumber string = string(isbn[numberIndex])

		// casts number to int to add to current bookSum
		number, err := cast.Atoi(currentNumber)
		// returns false if a letter or invalid value is encountered
		if err != nil {
			return false
		}

		value = number

		bookSum += (value * i)
		numberIndex += 1
	}

	// returns true if the sum is divisible by 11
	// else it returns false
	if bookSum%11 == 0 {
		return true
	} else {
		return false
	}
}

func get13Sum(isbn string) int {
	// returns the sum of ISBN-13
	var bookSum int = 0

	var numberIndex int = 0
	var multiply int

	// adds check number to books sum
	bookSum += getCheckNum(isbn)

	// itterates through 13 to 1
	// sets uneven values to 1 and even to 3 then multiplies them with the number to add to the bookSum
	for i := 13; i >= 2; i-- {

		if i%2 == 0 {
			multiply = 3
		} else {
			multiply = 1
		}
		// casts char byte value to string to get UTF-8 value
		var currentNumber string = string(isbn[numberIndex])
		// casts number to int to add to current bookSum
		number, _ := cast.Atoi(currentNumber)

		bookSum += (number * multiply)
		numberIndex += 1
	}
	return bookSum
}

func check13(isbn string) bool {
	// confirms if ISBN-13 string is valid or not
	// returns true for valid ISBN
	// returns false if ISBN is not valid
	var bookSum int = get13Sum(isbn)

	// returns true if the sum is divisible by 11
	// else it returns false
	if bookSum%10 == 0 {
		return true
	} else {
		return false
	}
}

func toISBN13(isbn string) string {
	// casts ISBN-10 to ISBN-13
	// returns string with new ISBN number
	var checkValue string
	var newISBN string = fmt.Sprintf("978" + isbn)

	// stores value of last digit to replace it if needed to convert into valid ISBN-13
	checkValue = string(isbn[9])
	checkValue = strings.ToLower(checkValue)

	// stores int value of checknumber
	checkNumber := getCheckNum(isbn)

	// confirms if new ISBN-13 is valid, if not the check digit is changed
	if check13(newISBN) {
		return newISBN

	} else {
		// stores total value of current ISBN
		// gets remainder that the total is over or under being divisible by 10
		var total int = get13Sum(newISBN)
		var over int = total % 10

		// checks if value is x and reduces value by the number the total is over the sum required to divide by 10
		if checkValue == "x" {
			// gets new value for check number and replaces x with it
			value := fmt.Sprintf("%v", checkNumber-over)
			newISBN = strings.ReplaceAll(strings.ToLower(newISBN), "x", value)

		} else {
			// returns true if the check number is greater than the remainder
			// removes the remainder from the check number and adds the new checknumber
			if checkNumber >= over {
				value := fmt.Sprintf("%v", checkNumber-over)
				// removes old checknumber
				newISBN = strings.TrimRight(newISBN, checkValue)
				newISBN = fmt.Sprintf("%v%v", newISBN, value)

				// returns true if the remainder is greater than the check number
				// the required total to reach a divisible number is then added to the check number
			} else if checkNumber <= over {
				// calculates the total needed to add to make the ISBN divisible by 10
				toAdd := 10 - over
				value := fmt.Sprintf("%v", checkNumber+toAdd)
				// removes old checknumber
				newISBN = strings.TrimRight(newISBN, checkValue)
				newISBN = fmt.Sprintf("%v%v", newISBN, value)
			}
		}
	}
	return newISBN
}

func isbnCheck(isbn string) {
	// checks if isbn is valid
	// prints invalid if it's neither a ISBN-10 or ISBN-13 value
	// converts ISBN-10 to ISBN-13
	// confirms if converted value is valid
	if checkType(isbn) == 13 {
		if check13(isbn) {
			fmt.Println("VALID")
		} else {
			fmt.Println("INVALID")
		}

	} else if checkType(isbn) == 10 {
		if check10(isbn) {

			isbn = toISBN13(isbn)
			isbnCheck(isbn)
			fmt.Println(isbn)

		} else {
			fmt.Println("INVALID")
		}

	} else {
		fmt.Println("INVALID")
	}
}

func main() {
	// stores isbn of user
	var isbn string

	var isRunning bool = true

	for isRunning {
		// requests isbn from user and then passes it to the variable isbn
		fmt.Println("\nPlease enter ISBN or type \"exit\" to quit")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			line := scanner.Text()
			isbn = line
			fmt.Printf("Input was: %q\n", line)
		}

		// removes all accidental white space in the ISBN number
		isbn = strings.ReplaceAll(isbn, " ", "")
		// removes all hyphens that might be included in official ISBN
		isbn = strings.ReplaceAll(isbn, "-", "")

		if strings.ToLower(isbn) == "exit" {
			isRunning = false
			return
		}

		// calls function to check isbn
		isbnCheck(isbn)

	}
}
