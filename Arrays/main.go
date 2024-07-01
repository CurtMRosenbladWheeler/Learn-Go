package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var a [5]int
	fmt.Printf("emp: %v\n", a)

	a[4] = 100
	fmt.Printf("Set a to %v\n", a[4])
	fmt.Printf("Get a at index 4 with value %v\n", a[4])
	fmt.Println(a)

	fmt.Println(len(a))

	var b [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("dcl: %v\n", b)

	var c = [...]int{1, 2, 3, 4, 5}
	fmt.Printf("dcl: %v\n", c)

	var d = [...]int{100, 3: 300, 400}
	fmt.Printf("idx: %v\n", d)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = rand.Intn(100)
		}
	}
	fmt.Println(twoD)

	var twoD2 = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(twoD2)
}
