package main

import "fmt"

func main() {

	//arrays

	//var fruitArr [2]string

	//assign values

	//fruitArr[0] = "Apple"
	//fruitArr[1] = "Orange"

	//fmt.Println(fruitArr)
	//fmt.Println(fruitArr[1])

	//slices
	fruitSlice := []string{"Apple", "Orange", "Grape"}

	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3])
}
