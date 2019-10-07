package main

import "fmt"

// Switch type
type Switch string

func (s *Switch) toggle() {
	if *s == "on" {
		*s = "off"
	} else {
		*s = "on"
	}
	fmt.Println(*s)
}

// Toggleable interface
type Toggleable interface {
	toggle()
}

func main() {
	s := Switch("off")
	var t Toggleable = &s
	t.toggle()
	t.toggle()
	t.toggle()
	t.toggle()
}
