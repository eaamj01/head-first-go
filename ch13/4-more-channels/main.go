package main

import (
	"fmt"
	"time"
)

func reportNap(name string, delay int) {
	for i := 0; i < delay; i++ {
		fmt.Println(name, "is sleeping")
		time.Sleep(1 * time.Second)
	}
	fmt.Println(name, "wakes up!")
}

func send(myChannel chan string) {
	reportNap("sending go routine", 2)
	fmt.Println("***Sending value***")
	// will block on this send while main is asleep
	myChannel <- "a"
	fmt.Println("***Sending value***")
	myChannel <- "b"
}

func main() {
	myChannel := make(chan string)
	go send(myChannel)
	reportNap("receiving go routine", 5)
	fmt.Println(<-myChannel)
	fmt.Println(<-myChannel)
}
