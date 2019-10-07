package main

import "fmt"

// Robot type
type Robot string

// MakeSound Robot
func (r Robot) MakeSound() {
	fmt.Println("Beep Boop")
}

//Walk Robot
func (r Robot) Walk() {
	fmt.Println("Powering legs")
}

// NoiseMaker interface
type NoiseMaker interface {
	MakeSound()
}

func main() {
	// define a variable with an interface type and assign a value of a type that satisfies the interface
	var noiseMaker NoiseMaker = Robot("Botco Ambler")
	// call a method that's part of the interface
	noiseMaker.MakeSound()
	// convert back to the concrete type using a type assertion
	var robot Robot = noiseMaker.(Robot)
	// call a method that's defined on the concrete type (not the interface)
	robot.Walk()
}
