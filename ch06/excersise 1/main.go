package main

import "fmt"

// Define a printLines function that takes a slice of strings as a parameter,
// and prints each element of that slice on a separate line.
// Then, in main, get a slice of the daysOfWeek array containing
// just the weekdays: “Monday”, “Tuesday”, “Wednesday”, “Thursday”, and “Friday”. Pass that slice to printLines.
func main() {
	daysOfWeek := [7]string{"Sunday", "Monday", "Tuesday",
		"Wednesday", "Thursday", "Friday", "Saturday"}

	weekDays := daysOfWeek[1:6]
	printLines(weekDays...)

}

func printLines(strings ...string) {
	for _, lineItem := range strings {
		fmt.Println(lineItem)
	}
}
