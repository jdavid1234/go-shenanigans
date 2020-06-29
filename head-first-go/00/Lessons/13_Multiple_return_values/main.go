package main

import "fmt"

//till now we have been used to passing in to the function 2+ arguments, but how about if you want the function to return 2+ values (or 1+ data type?

func addSubtract(a, b int) (int, int) {

	return a + b, a - b

}

func main() {

	addResult, subtractResult := addSubtract(45, 78)

	fmt.Println(addResult)
	fmt.Println(subtractResult)

	//if when we want to call our function, we do not need to add unecessary variables and then print them just for the sake of letting the program run.
	//lets say we want to call our addSubtract function (which returns 2 values) but only want the a+b logic and not the a-b logic, you can write like this:

	addOnlyResult, _ := addSubtract(607, 345)

	fmt.Println(addOnlyResult)

	//lets say we want to call our addSubtract function (which returns 2 values) but only want the a-b logic and not the a+b logic , you can write like this:

	_, subtrOnlyResult := addSubtract(64232, 5643)

	fmt.Println(subtrOnlyResult)

}
