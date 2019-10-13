package main

import (
	"fmt"

	"github.com/eaamj01/head-first-go/prose"
)

func main() {
	phrases := []string{"My parents", "a rodeo clown"}
	fmt.Println("A photo of", prose.JoinWithCommas(phrases))
	phrases = []string{"My parents", "a rodeo clown", "a prize bull"}
	fmt.Println("A photo of", prose.JoinWithCommas(phrases))
}
