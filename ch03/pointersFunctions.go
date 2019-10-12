package main

import "fmt"

func main() {
	// Assign the return pointer to a variable
	var myFloatPointer *float64 = createPointer()
	// Print the value at the pointer
	fmt.Println(*myFloatPointer)

	myBool := true
	printPointer(&myBool)

	amount := 6
	// pass a pointer instead of a variable
	double(&amount)
	fmt.Println(amount)
}

func createPointer() *float64 {
	myFloat := 98.5
	// return a pointer of sepecified the type
	return &myFloat
}

// use a pointer type for the parameter
func printPointer(myBoolPointer *bool) {
	// Print the value at the pointer that gets passed in
	fmt.Println(*myBoolPointer)
}

// Accept a pointer instaed of an int value
func double(number *int) {
	// Multiply the value at the pointer by 2
	*number *= 2
}
