package main

import "fmt"

func main() {
	var i uint = 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	for j := 0; j <= 5; j++ {
		fmt.Println(j)
	}

	for n := range 3 {
		fmt.Println("range", n)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 1000; n++ {
		if n%2 == 0 {
			fmt.Println("Even", n)
			continue
		} else {
			fmt.Println("odd", n)
			break
		}
	}
}
