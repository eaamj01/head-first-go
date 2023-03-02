package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func hello() {
	fmt.Println("hello")
	fmt.Println("Hello, \nGo!", "something else")

	// Takes a number and rounds it down
	fmt.Println(math.Floor(2.75))

	//Capitalises the first letter of each word
	fmt.Println(strings.Title("head first go"))

	// Rune literal - this prints 65
	fmt.Println('A')
	// Rune literal - this prints 9
	fmt.Println('\t')

	// int
	fmt.Println(reflect.TypeOf(42))

	// float64
	fmt.Println(reflect.TypeOf(3.145))

	// bool
	fmt.Println(reflect.TypeOf(true))

	// string
	fmt.Println(reflect.TypeOf("Hello, Go!"))

	var myInt = 2
	fmt.Println("myInt", reflect.TypeOf(myInt))
	// Type conversion from int to float64
	fmt.Println("myInt", reflect.TypeOf(float64(myInt)))

}
