package main

import "fmt"

func apples() {
	originalCount, eatenCount := 10, 4

	fmt.Println("I started with", originalCount, "apples")
	fmt.Println("Some jerk ate", eatenCount)
	fmt.Println("There are", originalCount-eatenCount, "apples left")
}
