package main

import (
	"fmt"
)

//Step 1: Define our Struct

type rectProps struct {
	width  float64
	height float64
}

//step 3: create a function and designate it to operate on our specific rectProps struct

func (rectangle rectProps) calcArea() float64 {

	fmt.Println("I've got it!")

	return rectangle.width * rectangle.height

}

func main() {

	//Step 2: assign specific numeric properties to our Struct into a variable
	recA := rectProps{width: 10.45, height: 5.62}
	recB := rectProps{width: 34.67, height: 56.24}
	recC := rectProps{width: 56.35, height: 67.35}
	recD := rectProps{width: 67.42, height: 15.68}

	//step 5: use our calcArea method to calculate the area of recA

	fmt.Println("the area of recA is:", recA.calcArea())
	fmt.Println("the area of recB is:", recB.calcArea())
	fmt.Println("the area of recC is:", recC.calcArea())
	fmt.Println("the area of recD is:", recD.calcArea())

}

//there is also a concept of using pointers in methods  - we'll get to it later
