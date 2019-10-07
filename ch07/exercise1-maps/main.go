package main

import "fmt"

func main() {
	// We'll count the number of times each letter occurs
	// within this slice.
	letters := []string{"b", "a", "c", "a", "b", "a",
		"a", "b", "c"}

	occurrences := map[string]int{}
	for _, letter := range letters {
		occurrences[letter]++
	}

	for letter, count := range occurrences {
		fmt.Println("letter", letter, "occurred", count)
	}
	// YOUR CODE HERE: Process each element of "letters",
	// keeping track of how many times you've seen "a",
	// "b", or "c".
	// Finally, print out the number of times each letter
	// occurred.
}
