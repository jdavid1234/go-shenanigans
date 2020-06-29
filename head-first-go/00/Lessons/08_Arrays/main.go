package main

import (
	"fmt"
)

func main() {

	var intArr [5]int
	fmt.Println(intArr)
	// this will give us an output of [0 0 0 0 0], since all we did is declare an array to store 5 intigers.
	// We did not say what these numbers are, so they default to 0

	var boolArr [10]bool
	fmt.Println(boolArr)

	// this will give us an output of [False False...10 times], since all we did is declare an array to store 5 Booleans.
	// We did not set them to True or Flase, so they default to False

	var boolStr [4]string
	fmt.Println(boolStr)

	//	This defaults to any array 4 blank spaces

	//We can later assign values to specific places within the array
	intArr[0] = 5
	intArr[1] = 10
	fmt.Println(intArr)

	//this will make the value of the first place 5, and the 2nd place 10

	//instead of printing a full array, we can print a specifc place within an array:
	fmt.Println(intArr[0])

	//	How to declare an array and assign values to it in short hand

	a := [5]int{1, 2, 3, 4, 5}

	fmt.Println(a)

	//how to print the length of an array as opposed to the arrat itself

	fmt.Println(len(a))

	//how to create multi dimentional arrays

	var aa [5][5]int
	fmt.Println(aa)

	//the first brackets contain the amount of arrays, and the 2nd brackets are how many units in each of them

	var bb [5][4]int
	fmt.Println(bb)

	/*
		a 2D array is another way of saying, columns and rows
		so  bb [5][4] is like saying, bb is a table with 5 rows and 4 columns
		and aa[5][5] is like saying, aa is a table with 5 rows and 5 columns
		all values within this table default to 0
		00000
		00000
		00000
		00000
		00000
		we can quickly assign values to this table using loops
		let's call each array set i, and each individual value within those arrays,

		i0
		00000

		i1

		00000

		i2
		00000

		i3

		00000

		i4

		00000

		----still do not fully grasp this
	*/
	count := 0

	for i := 0; i < 5; i++ {

		for j := 0; j < 5; j++ {

			aa[i][j] = count
			count = count + 1

			println(i, j, count)

		}

	}

	fmt.Println(aa)

	//string variables  - behind the scenes are really just arrays  -  what it is really doing is declaring an array with X amount of indexes (spaces)

	s := "hello, Ӽ"
	// hello := s[:5]
	// world := s[7:]

	// println(s)
	// println(hello)
	// println(world)

	fmt.Println("--------------------------")
	for i := 0; i < (len(s)); i++ {

		fmt.Printf(
			"%d    %q    \n",
			i,
			s[i],
		)

	}

	fmt.Println("--------------------------")
	for i, r := range "hi ӼӼ" {

		fmt.Printf(
			"%d   %q     %d     \n",
			i,
			r,
			r,
		)
	}

}
