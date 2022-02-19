package main

import "fmt"

func exercise1() {
	count := 20
	unitWeight := 0.4
	totalWeight := float64(count) * unitWeight
	fmt.Println(count, "cans weigh", totalWeight, "kilograms")
}
