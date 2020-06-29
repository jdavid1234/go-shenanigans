package main

import (
	"fmt"
)

func main() {

	//quick way to chnage the value of a variable:
	a := 5
	fmt.Println(a)
	a = 10
	fmt.Println(a)

	//although this chnaged the value of a, let's take  a look a what is going on behind the scenes:

	b := 7
	fmt.Println(b)
	fmt.Println(&b)

	b = 14
	fmt.Println(b)
	fmt.Println(&b)
	//you see here that we have now chnaged the value of b, from 7 to 14, but the place that this value of b is stored (0xc000018130) has not chnaged

	val := 20
	fmt.Println(val)

	var ptrVal *int = &val

	//this will print out the memory address where the value of 20 is stored
	fmt.Println(ptrVal)

	//this will print out the value stored at a specific memory address (dereferencing a pointer)
	fmt.Println(*ptrVal)

	//*ptrVal is currently 20, but what what would happen if we chnage this value?

	*ptrVal = 10
	fmt.Println(*ptrVal)

	// so now, the value we have stored inside the memory address (0xc000018108) has changed from 20 to 10

	//now although we have not explicity changed the value of 'val', since we have declared a change to the content stored within the memory address of val,
	// this has the power to change the value of val,
	fmt.Println(val)
	fmt.Println(ptrVal)

	//let's chnage the value of val, once more

	val = 50
	fmt.Println(val)
	//val has now chnaged from 10 to 50

	//and now, if we want to find the value of *ptrVal (the value stored at the memory address 0xc000018108)
	fmt.Println(*ptrVal)

	//so a change made to our variable (val), will be reflected inside *ptrVal as well

	///////////////////////////////////////
	// lets say we want to assign a function on a variable:

	/*


				Takeaway:
				x  = 10
						When you use &, this is the memory address of the variable you pass; a := &x
				and if you want to print the value (string or number) stored at a specific memory address, you use  * (dereferencing a pointer) , println(*a)


				side note: the * is also used when you want to declare the pointer variable
				so above, the proper way to say a := &x, is really
				var *int a = &x
		 k.0
		1k                                                                                                                  jmn	*/
}
