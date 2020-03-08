package main

import "fmt"

func main() {

	//valus no known in advance

	var myInts [3]int64

	myInts[1] = 99

	fmt.Println(myInts[0])
	fmt.Println(myInts[1])
	fmt.Println(myInts[2])

	//valus known in advance

	myInts2 := [3]int64{23, 0, 99999}

	fmt.Println(myInts2[0])
	fmt.Println(myInts2[1])
	fmt.Println(myInts2[2])

	//multi-lines

	myStrings := [5]string{

		"ayn",
		"od",
		"milvado",
		"efes",
		"zulaso",
	}

	fmt.Println(myStrings[0])
	fmt.Println(myStrings[1])
	fmt.Println(myStrings[2])
	fmt.Println(myStrings[3])
	fmt.Println(myStrings[4])

}
