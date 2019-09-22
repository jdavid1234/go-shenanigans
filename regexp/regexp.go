// B''H

package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {

	// -- ------------------------------------
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")

	fmt.Println("\n--------------------")
	fmt.Println("regexp.MatchString")
	fmt.Println(match)

	// -- ------------------------------------
	r, _ := regexp.Compile("p([a-z]+)ch")

	// -- ------------------------------------
	fmt.Println("\n--------------------")
	fmt.Println("r.MatchString")
	fmt.Println(r.MatchString("peach"))

	// -- ------------------------------------
	fmt.Println("\n--------------------")
	fmt.Println("r.FindString")
	fmt.Println(r.FindString("peach punch"))

	// -- ------------------------------------
	fmt.Println("\n--------------------")
	fmt.Println("r.FindStringIndex")
	fmt.Println(r.FindStringIndex("peach punch"))

	// -- ------------------------------------
	fmt.Println("\n--------------------")
	fmt.Println("r.FindStringSubmatch")
	fmt.Println(r.FindStringSubmatch("peach punch"))

	// -- ------------------------------------
	fmt.Println("\n--------------------")
	fmt.Println("r.FindStringSubmatchIndex")
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// -- ------------------------------------
	fmt.Println("\n--------------------")
	fmt.Println("r.FindAllString")
	fmt.Println(r.FindAllString("peach punch pinch", -1))

	// -- ------------------------------------
	fmt.Println("\n--------------------")
	fmt.Println("r.FindAllStringSubmatchIndex")
	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))

	// -- ------------------------------------
	fmt.Println("\n--------------------")
	fmt.Println("r.FindAllString")
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	// -- ------------------------------------
	fmt.Println("\n--------------------")
	fmt.Println("r.Match")
	fmt.Println(r.Match([]byte("peach")))

	// -- ------------------------------------
	fmt.Println("\n--------------------")
	fmt.Println("r.MustCompile")
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	// -- ------------------------------------
	fmt.Println("\n--------------------")
	fmt.Println("r.ReplaceAllString")
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// -- ------------------------------------
	fmt.Println("\n--------------------")
	fmt.Println("r.ReplaceAllFunc")

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
