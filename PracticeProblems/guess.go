//guess.go challenges players to guess a random number
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
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("I've chosen a random number between 1 and 100")
	fmt.Println("Can you guess it?")

	reader := bufio.NewReader(os.Stdin)

	success := false

	for guesses := 0; guesses < 10; guesses++ {

		fmt.Println("You have", 10-guesses, "guesses left.")

		fmt.Print("Make a guess: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("Oops, your guess was too LOW")
		} else if guess > target {
			fmt.Println("Oops, your guess was too HIGH")
		} else {
			success = true
			fmt.Println("You go it!  The number was", target)
			fmt.Println("You guessed correctly in ", guesses, "guesses")
			break
		}
	}

	if !success {
		fmt.Println("Sorry, you didn't guess my number.  It was:", target)
	}
}
