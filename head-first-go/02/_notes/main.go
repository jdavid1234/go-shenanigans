// B"H

package main

import "fmt"

func main() {

	// -- ---------------------------------
	// Slice example 1
	fmt.Println("------------------------")
	fmt.Println("Slice example 1:")
	fmt.Println("Use this 8% of the time")

	// -- ---------------------------------
	/*
		make:
			- built-in function
			- allocates and initializes an object of type slice, map, or chan (only).
			- the first argument is a type.
			- the return type is the same as the type of its argument, not a pointer to it.
			- The specification of the result depends on the type
	*/
	// -- ---------------------------------

	var notes []string
	notes = make([]string, 7)

	notes[2] = "hi"
	notes[3] = "how"
	notes[4] = "are"
	notes[5] = "you"

	for index, value := range notes {
		fmt.Println("index is", index, "- value is", value)
	}
	// -- ---------------------------------

	// -- ---------------------------------
	// Slice example 2
	fmt.Println("------------------------")
	fmt.Println("Slice example 2:")
	fmt.Println("Use this 90% of the time")

	names := make([]string, 7)

	names[2] = "baruch"
	names[3] = "yossi"
	names[4] = "mendy"

	for index, value := range names {
		fmt.Println("index is", index, "- value is", value)
	}
	// -- ---------------------------------

	// -- ---------------------------------
	// Slice example 3
	fmt.Println("------------------------")
	fmt.Println("Slice example 3:")
	fmt.Println("Slice Literals")
	fmt.Println("No need to use `make`")
	fmt.Println("Use this 2% of the time")
	fmt.Println("--                    --")

	cities := []string{"b7", "jlm", "tlv", "tzfat", "elad"}

	cities[1] = "yerushalayim"

	for index, value := range cities {
		fmt.Println("index is", index, "- value is", value)
	}
	// -- ---------------------------------

	// -- ---------------------------------
	// Slice example 4
	fmt.Println("------------------------")
	fmt.Println("Slice example 4:")
	fmt.Println("Sliced from array")
	fmt.Println("In general do not use - unless really needed.")
	fmt.Println("--                    --")

	citiesArray := [5]string{"b7", "jlm", "tlv", "tzfat", "elad"}

	// -- -    -    -    -    -    -    -
	fmt.Println("\n citiesSlice1:")

	citiesSlice1 := citiesArray[0:2]

	for index, value := range citiesSlice1 {
		fmt.Println("index is", index, "- value is", value)
	}

	// -- -    -    -    -    -    -    -
	fmt.Println("\n citiesSlice2:")

	citiesSlice2 := citiesArray[2:5]

	for index, value := range citiesSlice2 {
		fmt.Println("index is", index, "- value is", value)
	}

	// -- ---------------------------------
	fmt.Println("Make a change in citiesArray")
	citiesArray[0] = "beer sheva"

	for index, value := range citiesSlice1 {
		fmt.Println("index is", index, "- value is", value)
	}

	fmt.Println("-  -  -  -  -  -  -  -  ")

	for index, value := range citiesSlice2 {
		fmt.Println("index is", index, "- value is", value)
	}

	// -- ---------------------------------

	// -- ---------------------------------
	// General appropriate usage of slice
	fmt.Println("------------------------")

	shoppingList := make([]string, 7)

	fmt.Println("\n print 1:")
	fmt.Println(shoppingList)

	// -- -   -   -   -   -   -   -   -   -
	shoppingList[2] = "apples"
	shoppingList[4] = "bread"

	fmt.Println("\n print 2:")
	fmt.Println(shoppingList)

	// -- -   -   -   -   -   -   -   -   -
	fmt.Println("ALWAYS use same var name when appending!!!")
	shoppingList = append(shoppingList, "milk", "tomatoes", "wine")

	fmt.Println("\n print 3:")
	fmt.Println(shoppingList)

	// -- ---------------------------------

}
