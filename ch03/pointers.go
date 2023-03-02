package main

import (
	"fmt"
	"reflect"
)

func pointers() {
	fmt.Println("pointers")

	myInt := 4
	// prints *int
	fmt.Println(reflect.TypeOf(&myInt))
	fmt.Println(myInt)
	myIntPointer := &myInt

	// prints memory address
	fmt.Println(myIntPointer)

	// prints value of variable at the pointer (4)
	fmt.Println(*myIntPointer)

	*myIntPointer = 8

	// myInt is now 8
	fmt.Println(myInt)

	var myFloat float64
	fmt.Println(reflect.TypeOf(&myFloat))

	var myBool bool
	fmt.Println(reflect.TypeOf(&myBool))

}
