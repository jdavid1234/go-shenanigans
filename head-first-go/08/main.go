// B"H

package main

import (
	"fmt"

	"github.com/jdavid1234/go-shenanigans/head-first-go/08/magazine"
)

func main() {

	// -- -----------------------------------------------
	subscriber1 := magazine.Subscriber{Name: "Aman Singh"}

	fmt.Println(subscriber1)

	subscriber1.Rate = 50.35
	subscriber1.Active = true

	subscriber1.Street = "123 Oak St"
	subscriber1.City = "Omaha"
	subscriber1.State = "NE"
	subscriber1.PostalCode = "68111"

	fmt.Println(subscriber1)

	fmt.Println("Street:", subscriber1.Street)

	// -- -----------------------------------------------
	subscriber2 := magazine.Subscriber{Name: "Rasha B'galui"}

	fmt.Println(subscriber2)

	subscriber2.Rate = 99.87

	subscriber2.Street = "999 Main St"
	subscriber2.City = "Miami"
	subscriber2.State = "FL"

	fmt.Println(subscriber2)

	fmt.Println("Street:", subscriber2.Street)
	fmt.Println("City:", subscriber2.City)

	// -- -----------------------------------------------

	// -- -----------------------------------------------
	employee := magazine.Employee{Name: "Joy Carr"}

	employee.Street = "456 Elm St"
	employee.City = "Portland"
	employee.State = "OR"
	employee.PostalCode = "97222"

	fmt.Println("Street:", employee.Street)
	fmt.Println("City:", employee.City)
	fmt.Println("State:", employee.State)
	fmt.Println("Postal Code:", employee.PostalCode)

	// -- -----------------------------------------------
	addr := magazine.Address{Street: "174th stg"}

	addr.City = "Portland"
	addr.State = "OR"
	addr.PostalCode = "97222"

	fmt.Println("Street:", addr.Street)
	fmt.Println("City:", addr.City)
	fmt.Println("State:", addr.State)
	fmt.Println("Postal Code:", addr.PostalCode)
}
