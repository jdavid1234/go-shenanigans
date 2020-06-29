package main

import (
	"fmt"
)

func main() {

	i := 2

	switch i {

	case 1:
		fmt.Println("one")

	case 2:
		fmt.Println("two")

	case 3:

		fmt.Println("three")

	}

	/*since we have declared that i = 2, the switch statments then check which number case it matches.
		  Since i is 2, this matches case 2, and therefore prints the output of case 2

	If for example I would change i to 4, it will not reutrn an error but there will be no output, since there is no case that matches the number 4.

	*/

	/*it will not work if you try to make the cases with letters
	var a string = "A"

	switch a {
	case A:
		fmt.Println("the letter A")
	case B:
		fmt.Println("the letter B")
	case C:
		fmt.Println("the letter C")

	}
	*/

	//how to write multiple Cases in 1 switch statement:

	j := 3

	switch j {

	case 1, 2:
		//meaning if j is = to either 1 or 2, this switch statement will be triggered. and if so, it will print the following:

		fmt.Println("one or 2")

	case 3:
		fmt.Println("three")

	case 4, 5:

		fmt.Println("four or five")

	default:
		//this means that is j does not match any of the above case numbers, it will instead revert to default and print:
		fmt.Println("not 1, 2, 3, 4, or 5")

	}

	//using a switch statement without numbering the cases:

	b := 9
	fmt.Println(b)
	switch {
	case b == 3:
		fmt.Println("b is equal to 3")

	case b > 3:
		fmt.Println("b is more than 3")

	case b < 3:
		fmt.Println("b is less than 3")

	}

}
