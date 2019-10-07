// guess challenges a players to guess a random number
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
	fmt.Println("I've choesn a random number between 1 and 100.")
	fmt.Println("Can you guess it?")
	// fmt.Println(target)

	reader := bufio.NewReader(os.Stdin)
	success := false
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have", 10-guesses, "left")
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
			fmt.Println("Good Job! You guessed it in ", guesses)
			success = true
			break
		}
	}
	if !success {
		fmt.Println("Sorrym you didn't guess my number. It was:", target)
	}
}
