// B"H

package main

import (
	"fmt"
	"log"

	"github.com/Ylazerson/go-shenanigans/head-first-go/07/datafile"
)

// -- ------------------------------------------
// THIS IS OBVIOUSLY NOT THE CORRECT APPRAOCH TO USE
// ... but hey for now .....
const filePath = "/home/laz/tmp/ch07-votes.txt"

func main() {

	// -- --------------------------------------
	lines, err := datafile.GetStrings(filePath)

	if err != nil {
		log.Fatal(err)
	}
	// -- --------------------------------------

	// -- --------------------------------------
	counts := make(map[string]int)

	for _, line := range lines {
		counts[line]++
	}

	for name, count := range counts {
		fmt.Printf("Votes for %s: %d\n", name, count)
	}
	// -- --------------------------------------
}
