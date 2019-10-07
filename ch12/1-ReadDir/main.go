package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// get a slice full of values representing the contents of "my_directory"
	files, err := ioutil.ReadDir("my_directory")
	if err != nil {
		log.Fatal(err)
	}

	// for each file in the slice
	for _, file := range files {
		if file.IsDir() {
			fmt.Println("Directory:", file.Name())
		} else {
			fmt.Println("File:", file.Name())
		}
	}
}
