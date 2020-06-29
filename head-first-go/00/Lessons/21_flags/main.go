package main

import (
	"flag"
	"fmt"
)

var outputS *string = flag.String("outputS", "StringTest", "This is a test, to run string flags on on the command line")

//The data type of our variable OutputS is not String, but Pointer string.
//when this function is stored, it is not stored in the form of words or numbers,
//but it is rather stored in the form of an address
//therefore, when we print the variable outputS, an address is printed
// when we print *outputS, this is the underlying content out our pointer variable; and a word is printed

func main() {

	//we use the String function in the flag package.
	//we set a word to use in our command line to access our flag
	//we set a default output it will return for cases that we do not type anything.
	//So if we type -outputS "StringTest"  or just -outputS, it will retun "StringTest".
	//and if we type -outputS "hi there", it will return "hi there"
	//we then type a phrase to describe what this flag is, which will be shown when there is and errow when using the help menu
	//this function will now be contained in a variable called outputS

	flag.Parse()

	fmt.Println(outputS)
	fmt.Println(*outputS)

	var outputB *bool = flag.Bool("outputB", false, "This is a test, to run Bool flags on on the command line")

	fmt.Println(outputB)
	fmt.Println(*outputB)

	outputF := flag.Float64("outputF", 0, "This is a test, to run float flags on on the command line")
	fmt.Println(*outputF)
	fmt.Println(outputF)

	outputI := flag.Float64("outputI", 0, "This is a test, to run Int flags on on the command line")
	fmt.Println(*outputI)
	fmt.Println(outputI)

}
