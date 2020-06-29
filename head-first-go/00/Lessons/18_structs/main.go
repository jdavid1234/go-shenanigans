package main

import (
	"fmt"
)

//first step is to define the struct

type employee struct {
	firstName string
	lastName  string
	id        int
	smoker    bool
}

func main() {

	fmt.Println(employee{firstName: "Homer", lastName: "Simpson", id: 1, smoker: true})

	//or you can set the details of the struct in shorthand:

	fmt.Println(employee{"Whalen", "Smithers", 2, false})

	//if you leave some of the details blank when setting the struct details, these will be auto-set to default values:

	fmt.Println(employee{firstName: "Frank", lastName: "", id: 56, smoker: false})
	//actually it looks like if you leave out any details, it returns an error (too many commas)

	//instead of directly printing the struct, you can assign the values of the struct to a variable:

	empFour := employee{"Monty", "Burns", 234, true}

	//we can then print the entire variable,

	fmt.Println(empFour)

	//or we can extract specific fields of the variable

	fmt.Println(empFour.firstName)
	fmt.Println(empFour.lastName)
	fmt.Println(empFour.id)
	fmt.Println(empFour.smoker)

	//We can also assign a pointer to a Struct

	empFourPtr := &empFour
	// empFourPtr is now a reference which points to the empFour variable

	fmt.Println(empFourPtr)

	// We can thn use the Pointer to Output specific fields within the empFour variable

	fmt.Println(empFourPtr.firstName)

	// we can also update the field of any struct using a Pointer
	empFourPtr.firstName = "Marge"

	//this should now have updated 2 things at once:

	//1) the pointer of empFourPtr which points to fields of the empFour struct, has now chnaged the firstName field from "Monty" to "Marge"

	fmt.Println(empFourPtr.firstName)

	//2) In addition, the struct itself has now changed the firstName field from "Monty" to "Marge"

	fmt.Println(empFour.firstName)

}
