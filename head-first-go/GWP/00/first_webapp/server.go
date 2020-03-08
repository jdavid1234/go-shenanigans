package main

import "fmt"

func main() {

	ids := []int{12, 34, 56, 34, 56, 78}

	//loop thru ids

	for i, id := range ids {

		fmt.Printf("%d - ID: %d\n", i, id)

	}

	//not using index? make sure to use an _ in order to not to declare somthing and not use it
	for _, id := range ids {

		fmt.Printf("ID: %d\n", id)

	}

	//add ids together

	sum := 0

	for _, id := range ids {

		sum = sum + id

	}

	fmt.Println("sum", sum)

	//range with map

	emails := map[string]string{"Bob": "bob@gmail.com", "sharon": "sharon@gmail.com"}

	for k, v := range emails {

		fmt.Printf("%s: %s\n", k, v)

	}

}

/*package main

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {

	fmt.Fprintf(writer, "Hi there, %s!", request.URL.Path[1:])

}

func main() {

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

*/
