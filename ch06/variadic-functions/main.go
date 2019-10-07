package main

import (
	"fmt"
	"math"
)

// A Variadic function is defined with ... (ellipsis) before the type of the last function
// or only function parameter in the function list
func main() {
	// Println function can take as many arguments as needed
	fmt.Println(1)
	fmt.Println(1, 2, 3, 4, 5)

	letters := []string{"a"}
	// append can take 1 or many items to add
	letters = append(letters, "b")
	letters = append(letters, "c", "d", "e", "f", "g")

	severalInts(1)
	severalInts(1, 2, 3)

	intSlice := []int{1, 2, 3}
	// use an int slice instead of the variadic arguments
	severalInts(intSlice...)

	severalStrings("a", "b")
	severalStrings("a", "b", "c", "d", "e")
	// if no arguments an empty slice is received
	severalStrings()

	stringSlice := []string{"a", "b", "c", "d"}
	// use a string slice instead of the variadic arguments
	mix(1, true, stringSlice...)

	fmt.Println(maximum(71.8, 56.2, 89.5))
	fmt.Println(maximum(90.7, 89.7, 98.5, 92.3))

	fmt.Println(inRange(1, 100, -12.5, 3.2, 0, 50, 103.5))
	fmt.Println(inRange(-10, 10, 4.1, 12, -12, -5.2))

	fmt.Println(average(100, 50))
	fmt.Println(average(90.7, 89.7, 98.5, 92.3))

}

// The numbers variable will hold a slice with the arguments
func severalInts(numbers ...int) {
	fmt.Println(numbers)
}

func severalStrings(stringData ...string) {
	fmt.Println(stringData)
}

func maximum(numbers ...float64) float64 {
	// start with avery low value
	max := math.Inf(-1)
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max
}

func inRange(min float64, max float64, numbers ...float64) []float64 {
	var result []float64
	for _, number := range numbers {
		if number >= min && number <= max {
			result = append(result, number)
		}
	}
	return result
}

func average(numbers ...float64) float64 {
	sum := 0.0
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}

func mix(num int, flag bool, strings ...string) {
	fmt.Println(num, flag, strings)
}
