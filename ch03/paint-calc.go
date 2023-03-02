package main

import (
	"errors"
	"fmt"
)

var metersPerLitre float64

func paintCalc() {
	fmt.Println("paintCalc")
	err := errors.New("height can't be negative")
	fmt.Println(err.Error())

	var amount, total float64
	metersPerLitre = 10.0
	amount, err = paintNeeded(4.2, -3.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%.2f litres needed\n", amount)
	}

	total += amount

	amount, err = paintNeeded(5.2, 3.5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%.2f litres needed\n", amount)
	}
	total += amount

	amount, err = paintNeeded(5.0, 3.3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%.2f litres needed\n", amount)
	}
	total += amount

	fmt.Printf("Total: %.2f litres needed\n", total)
}

func paintNeeded(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", height)
	}
	area := width * height
	return area / metersPerLitre, nil
}
