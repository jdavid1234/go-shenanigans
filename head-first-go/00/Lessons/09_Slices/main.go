package main

import (
	"fmt"
)

func main() {

	s := make([]int, 3)

	//slices are similar to arrays

	s[0] = 1
	s[1] = 2
	s[2] = 3

	fmt.Println(s)

	fmt.Println(s[0])
	fmt.Println(len(s))

	// append functions is unique to slices

	s = append(s, 4)

	fmt.Println(s)

	fmt.Println(len(s))

	s = append(s, 5, 6, 7)

	fmt.Println(len(s))

	//slice syntax

	fmt.Println(s[1:3])
	fmt.Println(s[:3])
	fmt.Println(s[1:])

	//concise slice definition:
	/*
		t := []int{100, 200, 300}

		fmt.Println(t)

		x := s

		fmt.Println(x)

		x[0] = 500
		fmt.Println(x)
		fmt.Println(s)

		//use copy to prevent changing both x and s
	*/
	x := make([]int, len(s))
	//instead of saying, make this slice with 7 units, we say, make this slice the amount of units (length of) s (which is also 7)

	copy(x, s)
	//meaning I want to take the current content of s and copy it into x

	x[0] = 500
	fmt.Println(x)
	fmt.Println(s)
	//even though we changed x after we copied s into it, s still remains the same

	//2-D slice. Similar to arrays, although length of slice can vary
	ss := make([][]int, 3)

	fmt.Println(ss)

	//what if we want to use loops to auto-populate this 2-D slice and get output like this:
	//[[0][1,2][2,3,4]]
	/*dont fully understand this code either
	  for i := 0; 1 < 3; i++ {

	  innerLen := i +1

	  ss[i] make([]int, innerLen)

	  for j := 0; j <innerLen; j++ {

	  	ss[i][j = i + j]
	  }

	  }


	  fmt.Println(ss)

	*/

}
