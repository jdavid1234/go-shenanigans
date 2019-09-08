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

	pages := make(chan Page)

	urls := []string{
		"https://example.com/",
		"https://golang.org/",
		"https://golang.org/doc"
	}

	for _, url := range urls {
		go responseSize(url, pages)
	}

	for i := 0; i < len(urls); i++ {
		page := <-pages
		fmt.Printf("%s: %d\n", page.URL, page.Size)
	}
}
