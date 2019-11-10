package main

import (
	"fmt"
)

func main() {

	//var now time.Time = time.Now()

	// -- --------------------------------

	var grade int = 65
	var passingGrade = 60

	if grade >= passingGrade {

		fmt.Println("you passed")

		var passed bool = grade >= passingGrade
		fmt.Println(passed)

		if passed {

			fmt.Println("you passed")

		}

	}

	// -- -------------------tets-------test-----
}
