package main

import "fmt"

var a = "a"

func scope() {
	fmt.Println("scope")
	a = "a"
	b := "b"
	if true {
		c := "c"
		if true {
			d := "d"
			fmt.Println(a)
			fmt.Println(b)
			fmt.Println(c)
			fmt.Println(d)
		}
		fmt.Println(a)
		fmt.Println(b)
		fmt.Println(c)
		// fmt.Println(d) out of scope
	}
	fmt.Println(a)
	fmt.Println(b)
	// fmt.Println(c) out of scope
	// fmt.Println(d) out of scope

}
