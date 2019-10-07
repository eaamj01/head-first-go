package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/datafile"
)

func main() {
	arrNumbers, err := datafile.GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	sum := 0.0
	for _, number := range arrNumbers {
		sum += number
	}
	fmt.Println(sum)

	sampleCount := float64(len(arrNumbers))
	fmt.Printf("Average: %0.2f\n", sum/sampleCount)
}
