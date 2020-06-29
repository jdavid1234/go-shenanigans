package main

import (
	"fmt"
)

//make a fuction that can do this on its own:
//5! , 5x4x3x2x1

func fact(n int) int {

	if n <= 1 {

		return 1
	}
	return n * fact(n-1)

}

func main() {

	fmt.Println(fact(5))
	fmt.Println(fact(6))
	fmt.Println(fact(20))
	fmt.Println(fact(20))
}
