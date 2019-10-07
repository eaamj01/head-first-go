package shapes

import (
	"errors"
	"fmt"
)

// Rectangle Holds length and width
type Rectangle struct {
	length float64
	width  float64
}

// Length Get rectangle length
func (r *Rectangle) Length() float64 {
	return r.length
}

// SetLength Validate and set incoming length
func (r *Rectangle) SetLength(length float64) error {
	if length <= 0 {
		return errors.New("Invalid length")
	}
	r.length = length
	return nil
}

// Width Get rectangle length
func (r *Rectangle) Width() float64 {
	return r.width
}

// SetWidth Validate and set incoming width
func (r *Rectangle) SetWidth(width float64) error {
	if width <= 0 {
		return errors.New("Invalid width")
	}
	r.width = width
	return nil
}

// YOUR CODE HERE: Convert this function to a method on
// the "rectangle" type named "info".
func (r *Rectangle) info() {
	fmt.Println("Length:", r.length)
	fmt.Println("Width:", r.width)
}

// Convert this function to a method on the "rectangle" type.
func (r *Rectangle) makeSquare() {
	if r.length > r.width {
		r.length = r.width
	} else {
		r.width = r.length
	}
}
