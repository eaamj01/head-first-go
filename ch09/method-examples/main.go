package main

import "fmt"

// Number int type
type Number int

// MyType string
type MyType string

func main() {
	doNumberStuff()
	pointerStuff()
}

func doNumberStuff() {
	ten := Number(10)
	ten.Add(4)
	ten.Subtract(5)

	four := Number(4)
	four.Add(3)
	four.Subtract(2)

	numberFive := Number(5)
	fmt.Println("Original value of number:", numberFive)
	numberFive.Double()
	fmt.Println("numberFive value  after calling Double():", numberFive)

}

// Add Number to otherNumber
func (n Number) Add(otherNumber int) {
	fmt.Println(n, "plus", otherNumber, "is", int(n)+otherNumber)
}

// Subtract otherNumber from Number
func (n Number) Subtract(otherNumber int) {
	fmt.Println(n, "minus", otherNumber, "is", int(n)-otherNumber)
}

// Double Receive a pointer Number then double it's value
func (n *Number) Double() {
	*n *= 2
}

// this is a demonstartion of bad code, all type methods can take either value or pointer receivers
// you shouldn't mix the two like this
func pointerStuff() {
	value := MyType("a value")
	pointer := &value
	value.method()

	// Value automatically converted to a pointer
	value.pointerMethod()
	// value at pointer automatically retreived
	pointer.method()
	pointer.pointerMethod()

}

func (m MyType) method() {
	fmt.Println("Method with value receiver")
}

func (m *MyType) pointerMethod() {
	fmt.Println("Method with pointer receiver")
}
