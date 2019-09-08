// B"H

package main

import (
	"log"
	"net/http"
)

// -- ----------------------------------------
func viewHandler(writer http.ResponseWriter, request *http.Request) {

	message := []byte("Hello, web!")

	_, err := writer.Write(message)

	if err != nil {
		log.Fatal(err)
	}
}

// -- ----------------------------------------
func main() {

	http.HandleFunc("/hello", viewHandler)

	// Start up the web server:
	err := http.ListenAndServe("localhost:8080", nil)

	log.Fatal(err)
}
