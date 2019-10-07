package main

import "fmt"

// Appliance interface
type Appliance interface {
	TurnOn()
}

// Fan type
type Fan string

// TurnOn Fan
func (f Fan) TurnOn() {
	fmt.Println("Spinning")
}

// CoffeePot type
type CoffeePot string

// TurnOn CoffeePot
func (c CoffeePot) TurnOn() {
	fmt.Println("Powering Up")
}

// Brew Coffee
func (c CoffeePot) Brew() {
	fmt.Println("Heating Up")
}

func main() {
	var device Appliance
	device = Fan("Windco Breeze")
	device.TurnOn()
	device = CoffeePot("LuxBrew")
	device.TurnOn()
	// this won't work  as not defined in the interface
	// device.Brew()
}
