package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {

	// type time.Time
	now := time.Now()

	// type int
	year := now.Year()
	fmt.Println(year)

	// Full Month string
	fmt.Println(now.Month())

	broken := "G# R#cks!"
	// Create a replacer to swap # to o
	replacer := strings.NewReplacer("#", "o")

	// push broken through the replacer
	fixed := replacer.Replace(broken)
	fmt.Println(fixed)

	for x := 1; x <= 3; x++ {
		fmt.Println(x)
	}
	fmt.Println("--------------")
	for x := 3; x >= 1; x-- {
		fmt.Println(x)
	}
	fmt.Println("--------------")
	for x := 2; x <= 3; x++ {
		fmt.Println(x)
	}

	fmt.Println("--------------")
	// 1 3 5 7 9
	for x := 1; x <= 10; x += 2 {
		fmt.Println(x)
	}
}
