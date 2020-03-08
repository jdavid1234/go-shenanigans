package main

import "fmt"

func main() {

	x := 76
	y := 67

	//if else

	if x < y {

		fmt.Printf("%d is less than %d\n", x, y)
	} else {

		fmt.Printf("%d is less than %d\n", y, x)
	}

	//else if
	color := "green"
	if color == "red" {

		fmt.Println("color is red")
	} else if color == "blue" {

		fmt.Println("color is blue")
	} else {

		fmt.Println("color is NOT blue or red")

	}

	//switch

	switch color {
	case "red":
		fmt.Println("color is red")
	case "blue":
		fmt.Println("color is blue")

	default:
		fmt.Println("color is Not Blue or red")
	}

}
