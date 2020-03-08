package main

import "fmt"

func main() {

	fmt.Println("hi there")
	var name = "Joseph"
	var age int64 = 30
	var iq = 120.4

	fmt.Println("Hi, may name is", name)
	fmt.Println("I'm", age, "years old.")
	fmt.Println("I'm quite smart. I have an iq of", iq)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", iq)
}
