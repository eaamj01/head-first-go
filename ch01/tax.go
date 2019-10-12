package main

import "fmt"

func main() {
	// int
	price := 100
	fmt.Println("Price is", price, "dollars")

	// float64
	taxRate := 0.08
	tax := float64(price) * taxRate
	fmt.Println("Tax is", tax, "dollars")

	total := float64(price) + tax
	fmt.Println("Total cost is", total, "dollars")

	availableFunds := 120
	fmt.Println(availableFunds, "dollars available")
	fmt.Println("Within budget?", total <= float64(availableFunds))

}
