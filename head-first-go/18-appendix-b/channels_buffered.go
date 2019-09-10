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

6:39:11      START then wait 10 seconds

a 6:39:21    Waiting ready in size-2 buffer
b 6:39:21    Waiting ready in size-2 buffer
c 6:39:21    Waiting ready but blocking sending goroutine from continuing
d 6:39:24    3 seconds later it is passed down channel

6:39:24      END
*/
// -- -------------------------------------
func main() {

	fmt.Println(getTime(time.Now()))

	// Note the size of the buffer channel is 2
	channel := make(chan string, 2)

	go sendLetters(channel)

	time.Sleep(10 * time.Second)

	fmt.Println(<-channel, getTime(time.Now()))
	fmt.Println(<-channel, getTime(time.Now()))
	fmt.Println(<-channel, getTime(time.Now()))
	fmt.Println(<-channel, getTime(time.Now()))

	fmt.Println(getTime(time.Now()))
}
