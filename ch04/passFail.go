package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/keyboard"
)

// passFail reports whether a grade is passing or failing.
func passFail() {
	fmt.Println("passFail")
	fmt.Print("Enter a grade: ")
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}

	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println("A grade of", grade, "is", status)
}
