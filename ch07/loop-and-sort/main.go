package main

import (
	"fmt"
	"sort"
)

func main() {
	grades := map[string]float64{"Alma": 74.2, "Rohit": 86.5, "Carl": 59.7}
	var names []string
	for name, grade := range grades {
		names = append(names, name)
		fmt.Printf("%s has a grade of %0.1f%%\n", name, grade)
	}
	// Sort the slice alphabetically
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s has grade of %0.1f%%\n", name, grades[name])
	}
}
