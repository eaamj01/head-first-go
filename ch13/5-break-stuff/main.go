package main

import "fmt"

func greeting(myChannel chan string) {
	myChannel <- "hi"
}

func main() {
	myChannel := make(chan string)

	// fatal error: all goroutines are asleep - deadlock!
	//myChannel <- "hi from main"

	go greeting(myChannel)
	fmt.Println(<-myChannel)
}
