package main

import (
	"fmt"
)

func main() {

	// more revision is required on this, but the basic point here is to loop over a particular range of Arrays and Maps,
	//and then on every iteration of the loop, you make it do something, such a print a certain label next to each loop interation

	//You can range over key/value pairs in maps, and give a label to every kay and value
	m := map[string]int{"a": 1, "b": 1, "c": 3}
	for k, v := range m {

		fmt.Println("Key:", k, "Val:", v)
	}

	//if you want to label only the key and not the value:
	for k := range m {

		fmt.Println("Key:", k)
	}

	var listA []string = []string{"a", "b", "c"}

	fmt.Println(listA)

	for i, c := range listA {

		fmt.Println("index:", i, "content", c)

	}

}
