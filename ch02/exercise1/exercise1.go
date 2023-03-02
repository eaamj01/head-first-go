package main

import (
	"fmt"
	"strings"
)

func exercise1() {
	fmt.Println("exercise1")
	// The "builder" variable holds a "strings.Builder"
	// value. We're not assigning a value to "builder",
	// but that's okay; you can call methods on the
	// zero value for the "strings.Builder" type.
	var builder strings.Builder

	// YOUR CODE HERE: Call the "WriteRune" method on
	// "builder" 3 times, once with the rune 'a', once
	// with the rune 'b', and a third time with 'c'.
	builder.WriteRune('a')
	builder.WriteRune('b')
	builder.WriteRune('c')

	// YOUR CODE HERE: Call the "WriteString" method on
	// "builder" once, with the string "defg".
	builder.WriteString("defg")

	// YOUR CODE HERE: Call the "String" method on
	// "builder", and pass the return value to
	// "fmt.Println".
	fmt.Println(builder.String())
	return builder.String()
}
