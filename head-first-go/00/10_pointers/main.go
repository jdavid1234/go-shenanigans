package main

import "fmt"

func main() {

	a := 5
	b := &a
	c := &b
	fmt.Println(a, b, c)
	fmt.Printf("%T\n", b)

	// use * to read value from address in memory

	//the value inside b's bucket:
	fmt.Println(b)

	//the actual value that b is pointing to (so in a way, the * cancels out the &)
	fmt.Println(*b)

	//the value inside c's bucket, which is the adress in memory of where b is stored
	fmt.Println(c)

	//the actual value that c is pointing to
	//(in this case this is b, which is the memory address of a)
	fmt.Println(*c)

}
