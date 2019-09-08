// B"H

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// -- -----------------------------------------------
type Page struct {
	URL  string
	Size int
}

// -- -----------------------------------------------
// Get URL response then get the length of the response in bytes:
func responseSize(url string, channel chan Page) {

	fmt.Println("Getting", url)

	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	channel <- Page{URL: url, Size: len(body)}
}

// -- -----------------------------------------------
func main() {

	// -- -------------------------------------------
	// Create channel to carry a Page struct (see top).
	pages := make(chan Page)

	// -- -------------------------------------------
	/*
		1) Create a slice of string values with the URLs we want.
		2) Loop over the slice, and call responseSize with the current URL and the channel.
		3) Do a second, separate loop that runs once for each URL in the slice, and receives
		   and prints a value from the channel.

		It’s important to do this in a separate loop. If we received values in the same loop
		that starts the responseSize goroutines, the main goroutine would block until the
		receive completes, and we’d be back to requesting pages one at a time.
	*/
	// -- -------------------------------------------

	urls := []string{
		"https://example.com/",
		"https://golang.org/",
		"https://golang.org/doc"}

	for _, url := range urls {
		go responseSize(url, pages)
	}

	for i := 0; i < len(urls); i++ {

		page := <-pages

		fmt.Printf("%s: %d\n", page.URL, page.Size)
	}
}
