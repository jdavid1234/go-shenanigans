// B"H

package main

import (
	"log"
	"net/http"
)

// -- ----------------------------------------
func write(writer http.ResponseWriter, message string) {

	_, err := writer.Write([]byte(message))

	if err != nil {
		log.Fatal(err)
	}
}

// -- ----------------------------------------
func englishHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Hello, web!")
}

func frenchHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Salut web!")
}

func hindiHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Namaste, web!")
}

// -- ----------------------------------------
/*
➜  15b git:(master) ✗ go run salut.go

http://localhost:8080/hello
http://localhost:8080/salut
http://localhost:8080/namaste
*/
func main() {

	http.HandleFunc("/hello", englishHandler)
	http.HandleFunc("/salut", frenchHandler)
	http.HandleFunc("/namaste", hindiHandler)

	err := http.ListenAndServe("localhost:8080", nil)

	log.Fatal(err)
}
