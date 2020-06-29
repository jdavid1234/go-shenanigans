package main

import (
	"fmt"
	"math"
)

//step1) define our interface and input the methods which the interface will be comprised of:

type geometry interface {
	area() float64
	perim() float64
}

//step2) define our structs in order to establish which shapes will be used in this program

type rectangle struct {
	width  float64
	height float64
}

type circle struct {
	radius float64
}

//step 3) define our area methods which we will be using for each of our rectangle and circle structs

func (rec rectangle) area() float64 {

	return rec.width * rec.height

}

func (circ circle) area() float64 {

	return math.Pi * circ.radius * circ.radius
}

//step 4) define our perimeter methods which we will be using for each of our rectangle and circle structs

func (rec rectangle) perim() float64 {

	return (rec.width * 2) + (rec.height * 2)
}

func (circ circle) perim() float64 {

	return math.Pi * circ.radius * 2

}

//step 7: define a new method which will pass in our geometry interface
func measure(g geometry) {

	//the purpose of this method is to print 3 things:
	//when we later pass in our recA struct variable, the 2 fields of width and height will be passed into g
	fmt.Println(g)

	//our interface includes both our area() methods. So it will know to refer to either of our area methods.
	//if we pass recA into g, it will use the 'rectangle' struct to know to refer to the rectangle area() method
	//if we pass CircA into g, it will use the 'circle' struct to know to refer to the circle area() method

	fmt.Println(g.area(), "cm2")

	//our interface includes both our perim() methods. So it will know to refer to either of our perim methods.
	//if we pass recA into g, it will use the 'rectangle' struct to know to refer to the rectangle perim() method
	//if we pass CircA into g, it will use the 'circle' struct to know to refer to the circle perim() method

	fmt.Println(g.perim(), "cm")

}

func main() {

	//step 5: assign our rectangle struct to a variable, and assign values to each field
	//recA := rectangle{width: 3, height: 4}

	//step 6: assign our circle struct to a variable, and assign value radius field

	circA := circle{radius: 0}

	//measure(recA)
	measure(circA)

}
