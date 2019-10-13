package main

import (
	"fmt"
)

func a() {
	for i := 0; i < 50; i++ {
		fmt.Println("a")
	}
}

func b() {
	for i := 0; i < 50; i++ {
		fmt.Println("b")
	}
}

func greeting(myChannel chan string) {
	myChannel <- "hi"
}

func main() {
	// create a new channel
	myChannel := make(chan string)
	// pass the channel to a function running in a go routine
	go greeting(myChannel)
	//receive a value from the channel
	fmt.Println(<-myChannel)

	// go a()
	// go b()
	// // really poor way of ensuring main doesn't end before go routines
	// time.Sleep(1 * time.Second)
	fmt.Println("end main()")
}
