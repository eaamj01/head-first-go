package main

import (
	"fmt"
	"time"
)

// examples of array processings
func main() {
	// create an array of 5 elements that conatin ints
	// the elements are all assigned the value 0 until explicitly set
	var primes [5]int
	primes[0] = 2

	// prints 2
	fmt.Println(primes[0])

	// prints 0 as value not assigned
	fmt.Println(primes[2])

	// prints 0 as value not assigned
	fmt.Println(primes[4])
	fmt.Println(primes)

	var dates [3]time.Time
	dates[0] = time.Unix(1257894000, 0)
	dates[1] = time.Unix(1447920000, 0)
	dates[2] = time.Unix(1508632200, 0)

	fmt.Println(dates[1]) // 2nd element

	fmt.Println(dates) // 2nd element

	// strings get assigned as zero value empty string
	// until assign assigned explicitly
	var notes [7]string
	notes[0] = "do"

	// prints nothing as value not assigned
	fmt.Println(notes[2])

	// prints empty string as value not assigned
	fmt.Println(notes[4])
	fmt.Println(notes[0])
	fmt.Println(notes)

	// safe to increment array elements even if not explicitly set as they are
	// initialised to zero
	var counters [3]int
	counters[0]++
	counters[0]++
	counters[2]++
	fmt.Println(counters)

	// arrays can initialised using array literals
	arrNotes := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	fmt.Println(arrNotes)

	arrPrimes := [5]int{2, 3, 5, 7, 11}
	fmt.Println(arrPrimes)

	// multi line assignment

	arrText := [3]string{
		"This is a series of long strings",
		"which could be awkward to place",
		"together on a single line",
	}

	fmt.Println(arrText)

	// format arrays as they would appear in go code
	fmt.Printf("%#v\n", arrNotes)  // outputs [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	fmt.Printf("%#v\n", arrPrimes) // outputs [5]int{2, 3, 5, 7, 11}

	fmt.Println(len(arrNotes))

	// loop round arrNotes and print
	for i := 0; i < len(arrNotes); i++ {
		fmt.Println(arrNotes[i])
	}

	// safer way to loop an array
	for index, note := range arrNotes {
		fmt.Println(index, note)
	}

	// index not needed
	for _, note := range arrNotes {
		fmt.Println(note)
	}

	// only interested in unix
	for index := range arrNotes {
		fmt.Println(index)
	}
}
