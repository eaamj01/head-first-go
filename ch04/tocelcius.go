package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/keyboard"
)

func main() {
	fmt.Println("Enter a temperature in Fahrenheit: ")
	farenheit, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}

	celcius := (farenheit - 32) * 5 / 9
	fmt.Printf("%0.2f degrees Celcius\n", celcius)
}
