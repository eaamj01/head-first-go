package main

import "fmt"

// YOUR CODE HERE: Write a "perimeter" function.

func exercise2() {
	fmt.Println("exercise2")
	// YOUR CODE HERE: Call "perimeter" three times.
	// Add the three return values together, and store the
	// total in the "total" variable.

	total := perimeter(8.2, 10)
	total += perimeter(5, 5.4)
	total += perimeter(6.2, 4.5)
	fmt.Println("You'll need", total, "meters of fencing")
}

func perimeter(length, width float64) float64 {
	fencPerim := (length + width) * 2
	return fencPerim
}
