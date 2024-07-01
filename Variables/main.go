package main

import "fmt"

func main() {
	var mystring string = "Hello, world!"
	fmt.Println(mystring)

	var myint1, myint2 int = 1234, 5678
	fmt.Println(myint1)
	fmt.Println(myint2)

	var myemptyint int
	fmt.Println(myemptyint)

	var myint8 int8 = 88
	fmt.Println(myint8)

	var myint16 int16 = -1616
	fmt.Println(myint16)

	var myint32 uint32 = 323232 // uint does not accept negtive values
	fmt.Println(myint32)

	var myint64 int64 = 64646464
	fmt.Println(myint64)

	var mybool bool = true
	fmt.Println(mybool)

	var myfloat32 float32 = 123.4567890
	fmt.Println(myfloat32)

	var myfloat64 float64 = 123456.7890123456
	fmt.Println(myfloat64)

	myintvar := 16
	fmt.Println(myintvar)

	myfloatvar := 1234.567890
	fmt.Println(myfloatvar)

	myboolvar := false
	fmt.Println(myboolvar)
}
