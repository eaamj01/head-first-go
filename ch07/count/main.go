package main

import (
	"fmt"
	"log"

	"github.com/eaamj01/head-first-go/datafile"
)

func main() {
	lines, err := datafile.GetStrings("votes.txt")

	if err != nil {
		log.Fatal(err)
	}

	counts := make(map[string]int)

	for _, line := range lines {
		counts[line]++
	}
	fmt.Println(counts)
}
