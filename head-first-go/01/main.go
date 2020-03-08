// B"H

// Guess challenges players to guess a random number.

package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	// -- -----------------------------------------
	// Set the random number:
	var seconds int64 = time.Now().Unix()
	rand.Seed(seconds)
	var target int = rand.Intn(100) + 1
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")
	// -- -----------------------------------------

	// -- -----------------------------------------
	// Set up the reader to get the keyboard input:
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	// -- -----------------------------------------

	// -- -----------------------------------------
	// Play the game:
	var success bool = false

	// Start the loop:
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have", 10-guesses, "guesses left.")
		fmt.Print("Make a guess: ")

		// -- -------------------------------------
		// Get user input:
		input, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)

		// Convert alphanumeric string to integer (Atoi)

		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		// -- -------------------------------------

		// -- -------------------------------------
		// Check guess:
		if guess < target {
			fmt.Println("Oops. Your guess was LOW.")
		} else if guess > target {
			fmt.Println("Oops. Your guess was HIGH.")
		} else {
			success = true
			fmt.Println("Good job! You guessed it!")
			break
		}
		// -- -------------------------------------
	}

	if !success {
		fmt.Println("Sorry, you didn't guess my number. It was:", target)
	}
	// -- -----------------------------------------
}
