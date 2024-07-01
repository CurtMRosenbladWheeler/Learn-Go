package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n uint32 = 500000000

	const d float32 = 3e20 / float32(n)
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(float64(n)))
}
