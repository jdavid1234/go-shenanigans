package main

import (
	"fmt"

	"github.com/Ylazerson/go-shenanigans/stringutil"
)

func main() {

	// -- -------------------------------------
	fmt.Println(stringutil.Reverse("!oG ,olleH"))

	// -- -------------------------------------
	name := "Baruch"
	fmt.Println(stringutil.Reverse(name))

	// -- -------------------------------------
	var qty int
	qty = 1
	fmt.Println(qty)

	// -- -------------------------------------
	var qty2 int = 3
	fmt.Println(qty2)

	// -- -------------------------------------
	num := 12
	fmt.Println(num)

	num = 13
	fmt.Println(num)

	// -- -------------------------------------
	x, y := 15, 99
	fmt.Println(x)
	fmt.Println(y)

	// -- -------------------------------------
	var myInt int
	fmt.Println("myInt", myInt)

	var myFloat float64
	fmt.Println("myFloat", myFloat)

	var myBool bool
	fmt.Println("myBool", myBool)

	var myStr string
	fmt.Println("myStr", myStr)
}
