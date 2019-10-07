package main

import "fmt"

// Whistle type
type Whistle string

// MakeSound Whistle
func (w Whistle) MakeSound() {
	fmt.Println("Tweet!")
}

// Horn type
type Horn string

// MakeSound Horn
func (w Horn) MakeSound() {
	fmt.Println("Honk!")
}

// NoiseMaker interface
type NoiseMaker interface {
	MakeSound()
}

func play(n NoiseMaker) {
	n.MakeSound()
}

func main() {
	var toy NoiseMaker
	toy = Whistle("Toyco Canary")
	toy.MakeSound()

	toy = Horn("Toyco Blaster")
	toy.MakeSound()

	play(Whistle("Toyco Canary"))
	play(Horn("Toyco Blaster"))
}
