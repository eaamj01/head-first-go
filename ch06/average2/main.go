package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// take arguments from the command line and then convert them to a slice
// before passing them to the average() function for processing
func main() {
	// get item 2 onwards, in otherwords everything after "average2" programe name
	fmt.Println(os.Args[1:])
	arguments := os.Args[1:]
	var numbers []float64
	for _, argument := range arguments {
		number, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}

	fmt.Printf("Average: %0.2f\n", average(numbers...))
}

func average(numbers ...float64) float64 {
	sum := 0.0
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}
