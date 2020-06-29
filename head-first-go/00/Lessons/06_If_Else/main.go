package main

import "fmt"

func main() {

	i := 5

	if i%2 == 0 {

		fmt.Println(i, "is even.")

	} else {

		fmt.Println(i, "is odd.")

	}

	//this program will therefore output, "5 is odd"

	//We can also write an if else with a Loop:

	for j := 0; j <= 100; j++ {

		if j%2 == 0 {

			fmt.Println(j, "is even.")
		} else if j%2 == 0 {

			fmt.Println(j, "is divisible by 3.")
		} else if j%3 == 0 {

			fmt.Println(j, "is even AND divisible by 3.")

		} else {

			fmt.Println(j)
		}

	} // this curly brace ends our loop

}
