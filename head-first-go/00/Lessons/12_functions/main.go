package main

import (
	"fmt"
)

// we are going to define a function outside of our main program function

//we start by 'func' which declares the function
//then 'add', the name we are choosing for our function
//(a, b) these are the 2 argumnets we intend to pass in to the function
//(a int, b int) - we need to specify the type of variables that are going to be passed in as arguments to our function
// we need too add the return type for our function, in this case it would be int, (as we are using it to add 2 numbers together and have it equal a  number)
//inside the curly braces, we type the logic of our function

func add(a int, b int) int {

	return a + b
}

//when defining the function, if all arguments are of the same data type, you can just label them once:

func addFour(a, b, c, d int) int {

	return a + b + c + d
}

//here is a function that passes in and out 2 strings, a and b, and returns the same strings with an 'X' at the end of each string
func addX(a, b string) (string, string) {
	return a + "X", b + "X"
}

func main() {

	// our next step would be to call the function we declared above
	ansOne := add(1, 2)

	fmt.Println(ansOne)

	//let's now call our addFour function

	ansTwo := addFour(12, 45, 56, 48)

	fmt.Println(ansTwo)

	//let's now call our addX function

	wordOneX, wordTwoX := addX("bo", "lo")
	fmt.Println(wordOneX)
	fmt.Println(wordTwoX)
}
