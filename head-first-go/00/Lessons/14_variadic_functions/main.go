/*perfect example of a variadic function is the fmt.Println() function.

	fmt.Println("this", "is", "a", "variadic", "function")

	//here we are able to add any amount of string inputs to be taken into the function, and each of the arguments will be seperated by a single white space
      now lets call the mult function with as many values we need
*/
package main

import (
	"fmt"
)

//a variadic function is able to pass in any amount of arguments
//so for example, we want the function to pass in any amount of numbers, and have them all multiply by each other

//the 3 dots in front of the int means that there can be any amount of numbers passed into the function
func mult(numberList ...int) int {
	var answer int = 1

	//'answer' is the final answer of all our numbers multiplied together. (we need to start at 1 and not 0, otherwise our end answer will be 0 as well)

	for _, number := range numberList {
		answer = answer * number
	}
	return answer
}

//if you get rid of the for _,number and just write,
//for number := range numberList {
// it will think we are referring to the index of our first number
//and not our first number itself. So it will then take this 0 and multiply it by all the numbers and make our MultiAns at the end always = 0

//numberList is the list of numbers you'd pass into the function and 'number' is each number we are currently itterating rate through

func main() {

	MultiAns := (mult(1, 2, 3, 4, 5))

	fmt.Println(MultiAns)

}
