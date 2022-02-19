package main

import "fmt"

func tax(price int, taxRate float64) (float64, bool) {

	fmt.Println("Price is", price, "dollars")

	tax := float64(price) * taxRate
	fmt.Println("Tax is", tax, "dollars")

	total := float64(price) + tax
	fmt.Println("Total cost is", total, "dollars")

	availableFunds := 120
	fmt.Println(availableFunds, "dollars available")
	fmt.Println("Within budget?", total <= float64(availableFunds))
	return total, total <= float64(availableFunds)
}
