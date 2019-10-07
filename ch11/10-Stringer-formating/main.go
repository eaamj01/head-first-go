package main

import "fmt"

// Gallons type
type Gallons float64

// String format gallons
func (g Gallons) String() string {
	return fmt.Sprintf("%0.2f gal", g)
}

// Litres type
type Litres float64

// String format litres
func (l Litres) String() string {
	return fmt.Sprintf("%0.2f l", l)
}

// Milliltres type
type Milliltres float64

// String format milliltres
func (m Milliltres) String() string {
	return fmt.Sprintf("%0.2f mL", m)
}

func main() {
	fmt.Println(Gallons(12.09248342))
	fmt.Println(Litres(12.09248342))
	fmt.Println(Milliltres(12.09248342))
}
