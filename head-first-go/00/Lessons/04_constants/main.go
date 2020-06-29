package main

import (
	"fmt"
)

func main() {

	const pi float64 = 3.1415926

	fmt.Println(pi)

	const c int = 1000

	fmt.Println(c)

	/* this will not work. This we have already defined c as a constant which is 50
	   c = 50

	*/

	var d int = 1000

	fmt.Println(d)

	d = 500

	// This will work, since we are able to reassign a new value to a variable

	// There was also no need to write "var d in = 5000, since d has already been declared as an int"

	/*but what if I want to re-declare d as a new variable type?

		var d string = "happy"

		fmt.Println(d)

	This will not work either, since although we are able to chnage the value of d, we are not able to re-declare d as a new variable type

	*/

	//How about this
	var A bool = true

	fmt.Println(A)

	//We can now chnage the value of A:

	A = false
	fmt.Println(A)

	/*but if we try to chnage A to an int:

	var A int = 1000
	fmt.Println(A)

	 this will not work, as we are not able to re-declare d as a new variable type

	*/

}
