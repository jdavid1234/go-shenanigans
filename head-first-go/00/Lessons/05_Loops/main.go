package main

import (
	"fmt"
)

func main() {

	// there are a few ways you can use loops

	// 1) how to write out a loop the LONG way:

	//step 1: declare your starting point:

	i := 0

	//Step 2: declare the range of your loop, (ending point, 10)

	for i <= 10 {

		//Step 3: what would you like to do with this loop? Print it?

		fmt.Println(i)

		//Step 4: increment i by 1, so it will print every number between 1 and 10.
		i = i + 1

	}

	//new example, without all the comments:

	a := -6

	for a <= 16 {

		fmt.Println(a)

		a = a + 1

	}

	//This is how you write a loop in Short hand:

	for b := 7; b <= 19; b++ {

		// this is just another way of saying, we start at 7, we want to end at 19, and we are incrementing b by 1.

		fmt.Println(b)
		//but of course, now we will need to say what we want to do with the loop; print it.

	}

	//how to make a loop only print out content, if it meets a certain condition
	for k := 4; k <= 15; k++ {

		if k%2 == 0 {
			continue
			//this is another way of saying, Skip over the even numbers
		}

		fmt.Println(k)
	}

	/*the fizzbuzz problem:
	What if we want the loop to print the numbers from 1 to 100,
	except if the number is divisible by 3 then print 'fizz',
	if the number is divisible by 5 print 'buzz'
	or if the number if divisible by both print 'fizzbuzz'.
	*/

	for R := 0; R <= 100; R++ {

		if R%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if R%3 == 0 {
			fmt.Println("Fizz")
		} else if R%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(R)
		}

	}

	//How to write a never ending loop:

	for {
		//this will make the text keep printing forever
		fmt.Println("Keep Printing")

		//if you want it to stop, you need to type in, break - which will then give us only 1 sibgle output if 'keep printing'

		break
	}

}
