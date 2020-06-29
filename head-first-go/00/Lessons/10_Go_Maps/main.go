package main

import (
	"fmt"
)

func main() {

	//maps in Go are just like dictionaries in python
	/*they hold info in form of Entries, Keys, and values
	  e.g.

	  myMap1
	  employee1 = John
	  employee2 = Harry
	  employee3 = Bob

	  each listed employee is Key, each name is a value, and a pair of a key and value is an Entry
	*/

	m := make(map[string]int)

	//this means that we initialize a map, call it m, and the keys will be Strings, and the values will be intigers

	m["a"] = 0
	m["b"] = 1

	//if you want to print the entire map:
	fmt.Println(m)

	//if you want to print a specifc key within the map:
	fmt.Println(m["a"])

	//if you want to print the amount of Entries within the map:
	fmt.Println(len(m))

	//to delete a Key (and value) from the map:
	delete(m, "a")

	fmt.Println(m)

	//To initialize the map in shorthand:
	m2 := map[string]int{"key1": 1, "key2": 2}
	fmt.Println(m2)

	//To search for a specific value of a key within a map

	val, isValPresent := m["b"]
	fmt.Println(val)
	//this will return the int value of b
	fmt.Println(isValPresent)
	//this will return the Bool value of b (true or false)

	//if you are only interested in checking whether the key exists, and don't care what the value is

	_, isValPresent2 := m["a"]
	fmt.Println(isValPresent2)

	//the underscore above saves us the trouble of decalring a variable called val as well as printing it.
	//instead, its like we are saying, 'put the value aside for now, and just focus on whether the key exists'

}
