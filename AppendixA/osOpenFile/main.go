package main

import (
	"fmt"
	"log"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("aardvark.txt", options, os.FileMode(0600))
	check(err)
	_, err = file.Write([]byte("amazing!\n"))
	check(err)
	err = file.Close()

	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }
	// check(scanner.Err())

	// File's owner will have full permissions
	fmt.Println(os.FileMode(0700)) // File's owner will have full permissions
	fmt.Println(os.FileMode(0070)) // Users in the file group will have full permissions
	fmt.Println(os.FileMode(0007)) // All other users on the system have full permissions

	// print decimal and octal version
	for i := 0; i <= 19; i++ {
		fmt.Printf("%3d: %04o\n", i, i)
	}
}
