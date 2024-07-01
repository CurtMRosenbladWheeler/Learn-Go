package main

import (
	"fmt"
	"time"
)

func main() {
	const i uint8 = 4
	fmt.Printf("Write %v as ", i)
	switch i {
	case 1:
		fmt.Println("one.")
	case 2:
		fmt.Println("two.")
	case 3:
		fmt.Println("three.")
	case 4:
		fmt.Println("four.")
	case 5:
		fmt.Println("five.")
	default:
		fmt.Println("VALUE OUT OF BOUNDS.")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It is the weekend!")
	default:
		fmt.Println("It is a weekday!")
	}

	time := time.Now()
	switch {
	case time.Hour() < 12:
		fmt.Println("It is before noon.")
	default:
		fmt.Println("It is the afternoon.")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a boolean.")
		case int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64:
			fmt.Println("I'm an intager.")
		case string:
			fmt.Println("I'm a string.")
		default:
			fmt.Printf("Don't know the type of %v\n", t)
		}
	}

	whatAmI(1)
	whatAmI("Hej")
	whatAmI(true)
	whatAmI(2.3)
}
