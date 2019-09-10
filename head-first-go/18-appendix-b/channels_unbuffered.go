// B"H

package main

import (
	"fmt"
	"time"
)

// -- -------------------------------------
func sendLetters(channel chan string) {

	time.Sleep(3 * time.Second)
	channel <- "a"

	time.Sleep(3 * time.Second)
	channel <- "b"

	time.Sleep(3 * time.Second)
	channel <- "c"

	time.Sleep(3 * time.Second)
	channel <- "d"
}

// -- -------------------------------------
func getTime(dt time.Time) string {

	return fmt.Sprintf(
		"%d:%d:%d",
		dt.Hour(),
		dt.Minute(),
		dt.Second(),
	)
}

// -- -------------------------------------
/*
OUTPUT

6:43:23      START then wait 10 seconds

a 6:43:33    Waiting ready but blocking sending goroutine from continuing
b 6:43:36    3 seconds later ...
c 6:43:39    3 seconds later ...
d 6:43:42    3 seconds later ...

6:43:42      END

*/
// -- -------------------------------------
func main() {

	fmt.Println(getTime(time.Now()))

	// Note no buffer
	channel := make(chan string)

	go sendLetters(channel)

	time.Sleep(10 * time.Second)

	fmt.Println(<-channel, getTime(time.Now()))
	fmt.Println(<-channel, getTime(time.Now()))
	fmt.Println(<-channel, getTime(time.Now()))
	fmt.Println(<-channel, getTime(time.Now()))

	fmt.Println(getTime(time.Now()))
}
