package main

import (
	"fmt"
	"strconv"
)

//Define Person Struct

type person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

//Greeting method (value receiver)
func (p person) greet() string {

	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

//Has birthday method (pointer receiver)
func (p *person) hasBirthday() {
	p.age++
}

//Get married method (pointer receiver)
func (p *person) getMarried(spouseLastName string) {

	if p.gender == "m" {

		return
	} else {
		p.lastName = spouseLastName

	}

}

func main() {
	//intialize person using struct
	person1 := person{firstName: "bob", lastName: "patring", city: "Sydney", gender: "m", age: 25}
	person2 := person{firstName: "mary", lastName: "barginm", city: "Melbourne", gender: "f", age: 75}
	person3 := person{firstName: "sharon", lastName: "charter", city: "Canberra", gender: "f", age: 15}
	person4 := person{firstName: "david", lastName: "polski", city: "Adelaide", gender: "m", age: 48}
	person5 := person{firstName: "tenilia", lastName: "gresten", city: "Perth", gender: "f", age: 25}
	person6 := person{firstName: "celia", lastName: "krilondem", city: "Hobart", gender: "f", age: 65}
	person7 := person{firstName: "glomb", lastName: "rochester", city: "Darwin", gender: "m", age: 27}

	fmt.Println(person1, person2, person3, person4, person5, person6, person7)
	person7.age++
	fmt.Println(person7)
	person1.hasBirthday()
	person1.hasBirthday()
	person1.hasBirthday()
	person1.hasBirthday()
	fmt.Println(person1.greet())
	fmt.Println(person3.lastName)
	fmt.Println(" ")
	person2.getMarried("Kanrokitov")
	fmt.Println(person2.greet())
	fmt.Println(" ")
	person4.getMarried("Gavrinken")
	fmt.Println(person4.greet())
}
