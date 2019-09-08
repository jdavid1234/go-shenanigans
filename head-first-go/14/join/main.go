// B"H

package main

import (
	"fmt"

	"github.com/Ylazerson/go-shenanigans/head-first-go/14/prose"
)

func main() {

	phrases := []string{"my parents", "a rodeo clown"}
	fmt.Println("A photo of", prose.JoinWithCommas(phrases))

	phrases = []string{"my parents", "a rodeo clown", "a prize bull"}
	fmt.Println("A photo of", prose.JoinWithCommas(phrases))
}
