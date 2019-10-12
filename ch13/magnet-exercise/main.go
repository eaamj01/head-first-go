package main

func main() {

}

func odd(channel chan int) {
	channel <- 1
	channel <- 3
}

func even(channel chan int) {
	channel <- 2
	channel <- 4
}
