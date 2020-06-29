package main

import (
	"fmt"
)
​
/*
a function has a command with parenthesis after it
*/
​
// A function can receive values and return values
//this kind of function will typically have 6 components                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            
/*i)the function name: receiveReturn
 ii) the wildcard variable; this variable will be a copy of the value (from the variable passed in from our func main): valueReceived 
 iii) the data type of our wildcard variable: string 
iv) the data type which will be passed out; string 
    // when the function does any sort of action on the data we pass in, it needs to know in which data type to pass out the info  
v) everything within the {     } is the action the fuction wll peform. this can include regular println functions like "this is the value I received"
vi) it can also include a 'return', which will do something directly connected with the values being passed in, (valueReceived) 
      //and pass it out with the pass-out data type (string)

 */ 
func receiveReturn(valueReceived string) string {
​
	
​
	fmt.Println("This is the value I received:")
	fmt.Println(valueReceived)
​
	return "Value Returned"
​
}
​
// A function doesn't have to receive or return. It can also do one or the other
func noReceiveNoReturn() {
	fmt.Println("I received nothing and returned nothing")
}
​
func main() {
	var valuePassedIn string = "Here I am"
	var valueReturned string
​
	fmt.Println("Here are the values before they are passed in and received:")
	fmt.Println("The value of the variable to be passed in:", valuePassedIn)
	fmt.Println("The value of the variable that will be returned:", valueReturned)
​
	// -- ----------------------------------------------------------
	fmt.Println("---------")
​
	// The function that received a passed-in value works like this
	valueReturned = receiveReturn(valuePassedIn)
​
	fmt.Println("\nHere is the value the function returned:")
	fmt.Println(valueReturned)
​
	// -- -----------------------------------------------------------
	fmt.Println("---------")
​
	// The function that neither receives nor returns works like this
	noReceiveNoReturn()
​
}