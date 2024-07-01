package main

import (
	"fmt"
)

func main() {
	const num1 uint8 = 7
	const num2 uint8 = 8

	if num1%2 == 0 {
		fmt.Printf("The number %v is even.\n", num1)
	} else {
		fmt.Printf("The number %v is odd.\n", num1)
	}

	if num1%2 == 0 || num2%2 == 0 {
		fmt.Printf("Either %v or %v are even.\n", num1, num2)
	}

	if num1%2 == 0 && num2%2 == 0 {
		fmt.Printf("Both %v or %v are even.\n", num1, num2)
	} else {
		fmt.Printf("Either %v or %v are odd.\n", num1, num2)
	}

	if num := 10; num > 10 {
		fmt.Printf("%v is greater than 10.\n", num)
	} else if num < 10 {
		fmt.Printf("%v is less than 10.\n", num)
	} else {
		fmt.Printf("%v is 10.\n", num)
	}
}
