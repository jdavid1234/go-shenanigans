// B"H

package main

import (
	"fmt"
	"time"
)

/*
goroutines send their values over their channels so quickly that
it’s hard to see what’s going on.

Here’s another program that slows things down so you can see the blocking happen.
*/

// -- ---------------------------------------
func reportNap(name string, delay int) {

	for i := 0; i < delay; i++ {

		fmt.Println(name, "sleeping")

		time.Sleep(1 * time.Second)
	}

	fmt.Println(name, "wakes up!")
}

// -- ---------------------------------------
func send(myChannel chan string) {

	reportNap("sending goroutine", 2)

	fmt.Println("***sending value 'a' down myChannel***")
	myChannel <- "a"

	fmt.Println("***sending value 'b' down myChannel***")
	myChannel <- "b"
}

// -- ---------------------------------------
func main() {

	myChannel := make(chan string)

	go send(myChannel)

	reportNap("receiving goroutine", 5)

	fmt.Println(<-myChannel)
	fmt.Println(<-myChannel)
}
