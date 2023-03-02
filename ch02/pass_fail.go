//pass fail reports whether  a grade is passing or failing.package first

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func passFail() {
	fmt.Println("passFail")
	fmt.Print("Enter a grade: ")

	// Set-up a bufferd reader that gets text from the keyboard
	reader := bufio.NewReader(os.Stdin)

	// Returns everything the user has typed, up to where  the newline rune will be read (hit Enter key)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(input)

	// Remove surrounding spaces and newline/carriage return
	input = strings.TrimSpace(input)

	// input is still a string so convert to float64
	grade, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}

	status := "failing"
	if grade >= 60 {
		status = "passing"
	}

	fmt.Println("A grade of", grade, "is", status)

	fileInfo, err := os.Stat("my.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("file size", fileInfo.Size())
}
