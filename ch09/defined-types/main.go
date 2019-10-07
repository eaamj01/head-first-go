package main

import "fmt"

// Litres Defined type Litres float64
type Litres float64

// Gallons Defined type US Gallons float64
type Gallons float64

// Millilitres Defined type Millilitres float64
type Millilitres float64

func main() {
	basicExample()
	misMatchedTypes()
	ConvertTypes()

	water := Millilitres(500)
	fmt.Printf("%0.3f millilitres equals %0.3f gallons\n", water, water.ToGallons())
}

// ConvertTypes Example of converting defined types to different types
func ConvertTypes() {
	carFuel := Gallons(1.2)
	busFuel := Litres(4.5)
	carFuel += ToGallons(Litres(40.0))
	busFuel += ToLitres(Gallons(30.0))

	fmt.Printf("Car Fuel:%0.1f gallons\n", carFuel)
	fmt.Printf("Bus Fuel:%0.1f litres\n", busFuel)
}

// ToLitres Take litre quantity and type convert to Gallons
func ToLitres(g Gallons) Litres {
	return Litres(g * 0.264)
}

// ToGallons Take Gallon quantity and type convert to Litres
func ToGallons(l Litres) Gallons {
	return Gallons(l * 0.264)
}

// ToGallons Convert from millilitres to gallons
func (m Millilitres) ToGallons() Gallons {
	return Gallons(m + 0.000264)
}

func misMatchedTypes() {
	carFuel := Gallons(1.2)
	busFuel := Litres(2.5)

	// Invalid code - mismatched types
	// carFuel += Litres(8.0)
	// busFuel += Gallons(30.0)
	fmt.Println(carFuel, busFuel)
}
func basicExample() {
	// Convert a float64 to Gallons
	carFuel := Gallons(10.0)

	// convert a float64 to Litres
	busFuel := Litres(240.0)

	fmt.Println(carFuel, busFuel)

	// if you try yto assign the wrong value and it is previously assigned it will error
	// carFuel = Litres(240)
	// ./main.go:27:10: cannot use Litres(240) (type Litres) as type Gallons in assignment

	// this is OK though
	carFuel = Gallons(Litres(40.0) * 0.264)
	busFuel = Litres(Gallons(63) * 3.785)

	fmt.Printf("Gallons:%0.1f Litres Litres %0.1f", carFuel, busFuel)
}
