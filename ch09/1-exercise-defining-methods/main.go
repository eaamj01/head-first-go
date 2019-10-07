package main

import "fmt"

type rectangle struct {
	length float64
	width  float64
}

// YOUR CODE HERE: Convert this function to a method on
// the "rectangle" type named "info".
func (r *rectangle) info() {
	fmt.Println("Length:", r.length)
	fmt.Println("Width:", r.width)
}

// Convert this function to a method on the "rectangle" type.
func (r *rectangle) makeSquare() {
	if r.length > r.width {
		r.length = r.width
	} else {
		r.width = r.length
	}
}

func main() {
	var r rectangle
	r.length = 4.2
	r.width = 2.3
	// YOUR CODE HERE: Update this function call to a
	// method call.
	r.info()
	r.makeSquare()
	r.info()
}
