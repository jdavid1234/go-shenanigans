package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"math"
)
​//step 1)
// we need to start out by declaring the following variables:
//there is a function in the flag package called float64.
//it allows you to type in a letter in the command line, assin your input a default value, as well a decription of what it does 
//since this is a float64 function, it views all cmd line input as numbers (floats)
//the letter for our width input will be w, it will have a default value of 0, and it means 'width of a rectangle'
//this will then be put into a variable called *width. Not width, but the address in memory where width is stored 


var width *float64 = flag.Float64("w", 0, "width of a rectangle")
var length *float64 = flag.Float64("l", 0, "length of a rectangle")
var radius *float64 = flag.Float64("r", 0, "radius of a circle")
​

//step 2) 
// we also meed to establish our potential error variables.
//the error package has a function called New, which allows us to customise our own error meesages, to be displayed to the user when the error is triggered 
//as we designate our custom errors, we put each of them into a variable (which have the data type of error)
var errChoose error = errors.New("cannot calculate rectangle and circle")
var errNeg error = errors.New("only positive and non-zero numbers")
​


//step3) define our interface and input the methods which the interface will be comprised of:
type geometry interface {
	area() float64
	perim() float64
}
​
//step4) define our structs in order to establish which shapes will be used in this program
type rectangle struct {
	width  float64
	length float64
}
type circle struct {
	radius float64
}
​
//step 5) define our area methods which we will be using for each of our rectangle and circle structs
func (rec rectangle) area() float64 {
	return rec.width * rec.length
}
func (circ circle) area() float64 {
	return math.Pi * circ.radius * circ.radius
}
​
//step 6) define our perimeter methods which we will be using for each of our rectangle and circle structs
func (rec rectangle) perim() float64 {
	return (rec.width * 2) + (rec.length * 2)
}
func (circ circle) perim() float64 {
	return math.Pi * circ.radius * 2
}
​
//step 7) define a new function which will pass in our geometry interface. (it is not a method since it is not assigned to operate on a specifc variable)
func measure(g geometry) {
	//the purpose of this method is to print 3 things:
	//when we later pass in our rec struct variable, the 2 fields of width and length will be passed into g
	//we will also pass in our circ struct variable, the 1 field of radius will be passed into g
	fmt.Println(g)
	//our interface includes both our area() methods. So it will know to refer to either of our area methods.
	//when we pass rec into g, it will use the 'rectangle' struct to know to refer to the rectangle area() method
	//when we pass Circ into g, it will use the 'circle' struct to know to refer to the circle area() method
	//so when we say g.area, it means that when rec iass passed into g, the rectangle area method is being performed on it, and will return a value of w x l     
	//we use printf, and %.2f means the float will be displayed to 2 decimal places. The \n just prints it on a new line
	fmt.Printf("Area is %.2f cm2\n", g.area())
	
	
	//our interface includes both our perim() methods. So it will know to refer to either of our perim methods.
	//when we pass rec into g, it will use the 'rectangle' struct to know to refer to the rectangle perim() method
	//when we pass Circ into g, it will use the 'circle' struct to know to refer to the circle perim() method
	fmt.Printf("Perimeter is %.2f cm\n", g.perim())
}
func main() {
​
	// parse the arguments. This activates the 3 pointer float varibales (flags) we defined above
	flag.Parse()

	​//we are able to declare our circ and rec struct variables now, but we do not want to assign values as this will be based on our cmd line input
	var rec rectangle
	var circ circle
​
	// make sure the user didn't mix rectangles and circles, handle arguments

	// we are tying to say here that if the user types a value for radius as well as width OR length as well as width, it shouold return an error and stop the program
	if (*length != 0 && *radius != 0) || (*width != 0 && *radius != 0) {
​
		log.Fatal(errChoose)
​

// we are tying to say here that if the user types a negative value for L, W, or R, this should return an error  
	} else if *width < 0 || *length < 0 || *radius < 0 {
​
		log.Fatal(errNeg)
​
// we are tying to say here that if the user types a 0 value for ALL L, W, and R, this should return an error  
	} else if *width == 0 && *length == 0 && *radius == 0 {
​
		log.Fatal(errNeg)
​
	} else if *width > 0 && *length > 0 {
​
		rec.width = *width
		rec.length = *length
		measure(rec)
​
	} else if *radius > 0 {
​
		circ.radius = *radius
		measure(circ)
​
	}
}









