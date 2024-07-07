package main

import (
	f "fmt"
)

func main() {
	var intSlice []int32 = []int32{4, 5, 6}
	f.Printf("The length is %v with capacity %v.\n", len(intSlice), cap(intSlice))
	f.Println(intSlice)

	intSlice = append(intSlice, 7)
	f.Printf("The length is %v with capacity %v.\n", len(intSlice), cap(intSlice))
	f.Println(intSlice)

	f.Println(intSlice[0:3])
	f.Println(&intSlice[3])

	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)
	f.Printf("The length is %v with capacity %v.\n", len(intSlice), cap(intSlice))
	f.Println(intSlice)
}
