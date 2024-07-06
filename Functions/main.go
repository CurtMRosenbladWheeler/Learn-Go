package main

import (
	"errors"
	f "fmt"
)

func main() {
	myString := "Hello, world!"
	printString(myString, true)

	numerator := 17
	denominator := 0
	var result2, remainder, err = intDivisionWRemainder(numerator, denominator)
	if err != nil {
		f.Println(err.Error())
	} else if remainder == 0 {
		f.Printf("The result of %v divided by %v is %v.", numerator, denominator, result2)
	} else {
		f.Printf("The result of %v divided by %v is %v with a remainder of %v.", numerator, denominator, result2, remainder)
	}

}

func printString(printValue string, newLine bool) {
	if newLine {
		f.Println(printValue)
	} else {
		f.Print(printValue)
	}
}

func intDivisionWRemainder(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}

	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
