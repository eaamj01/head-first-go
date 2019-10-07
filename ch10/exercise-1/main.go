package main

import (
	"fmt"
	"log"

	"github.com/eaamj01/head-first-go/shapes"
)

// check logs an error and exits if the error is not nil.
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var r shapes.Rectangle
	err := r.SetLength(4.2)
	check(err)
	fmt.Println("Length:", r.Length())
	// Set width to an invalid value.
	err = r.SetWidth(-2.3)
	check(err)
	fmt.Println("Width:", r.Width())
}
