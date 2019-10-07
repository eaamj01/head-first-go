package main

import (
	"fmt"
	"log"
)

// ComedyError type that holds a string
type ComedyError string

func (c ComedyError) Error() string {
	// The error needs to return a string so do a type conversion
	return string(c)
}

// OverHeatError type
type OverHeatError float64

func (o OverHeatError) Error() string {
	return fmt.Sprintf("Overheating by %0.2f degrees!", o)
}

func checkTemperature(actual float64, safe float64) error {
	excess := actual - safe
	if excess > 0 {
		return OverHeatError(excess)
	}
	return nil
}

func main() {
	// set up a variable with type error
	var err error
	err = ComedyError("What's a programmer's favourite beer? Logger!")
	fmt.Println(err)

	var err2 error = checkTemperature(121.379, 100)
	if err2 != nil {
		log.Fatal(err2)
	}
}
