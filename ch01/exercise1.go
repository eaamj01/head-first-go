package main

import "fmt"

func main() {
	count := 20
	unitWeight := 0.4
	totalWeight := float64(count) * unitWeight
	fmt.Println(count, "cans weigh", totalWeight, "kilograms")
}
