// B"H

/*
pass_fail reports whether a grade is passing or failing.
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	// ----------------------------------------
	// Get user input:
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	//Fatal function: log a message to the terminal and stop the program.
	if err != nil {
		log.Fatal(err)
	}
	// ----------------------------------------

	// ----------------------------------------
	// Cleanup input number:
	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}
	// ----------------------------------------

	// ----------------------------------------
	// See if it is a passing grade:
	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}

	fmt.Println("A grade of", grade, "is", status)
	// ----------------------------------------
}
