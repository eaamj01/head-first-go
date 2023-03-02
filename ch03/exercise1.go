package main

import "fmt"

// YOUR CODE HERE: Write a scoreSummary function that will
// work with the calls below.

func exercise1() {
	fmt.Println("exercise1")
	fmt.Printf("%10s | %8s | %8s | %8s | %8s\n",
		"Name", "Score 1", "Score 2", "Score 3", "Average")
	fmt.Println("------------------------------------------------------")
	scoreSummary("Jermaine", 95.4, 82.3, 74.6)
	scoreSummary("Jessie", 79.3, 99.1, 82.5)
	scoreSummary("Lamar", 82.2, 95.4, 77.6)
}

func scoreSummary(name string, score1 float64, score2 float64, score3 float64) {
	average := (score1 + score2 + score3) / 3.0

	fmt.Printf("%10s | %8.2f | %8.2f | %8.2f | %8.2f\n",
		name, score1, score2, score3, average)
}
