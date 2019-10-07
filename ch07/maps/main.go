package main

import "fmt"

func main() {
	// declare a map variable
	var ranks1 map[string]int
	// actually make the map
	ranks1 = make(map[string]int)
	ranks1["gold"] = 1

	// short way create a map and declare a variable to hold it
	ranks := make(map[string]int)

	ranks["gold"] = 1
	ranks["silver"] = 2
	ranks["bronze"] = 3
	fmt.Println(ranks["bronze"])

	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["Li"] = "Lithium"
	fmt.Println(elements["Li"])
	fmt.Println(elements["H"])

	isPrime := make(map[int]bool)
	isPrime[4] = false
	isPrime[7] = true
	fmt.Println(isPrime[4])
	fmt.Println(isPrime[7])

	// map literal
	ranks2 := map[string]int{"gold": 1, "silver": 2, "bronze": 3}
	fmt.Println(ranks2["gold"])
	fmt.Println(ranks2["bronze"])

	// multiline
	elements2 := map[string]string{
		"H":  "Hydrogen",
		"Li": "Lithium",
	}
	fmt.Println(elements2["Li"])
	fmt.Println(elements2["H"])

	// empty map
	emptyMap := map[string]float64{}
	println(emptyMap)

	//zero values within maps
	numbers := make(map[string]int)
	numbers["I've been assigned"] = 12
	fmt.Printf("%#v\n", numbers["I've been assigned"])

	// This will print 0
	fmt.Printf("%#v\n", numbers["I haven't been assigned"])

	//zero values within maps
	words := make(map[string]string)
	words["I've been assigned"] = "hi"
	fmt.Printf("%#v\n", words["I've been assigned"])

	// This will print an empty string ""
	fmt.Printf("%#v\n", words["I haven't been assigned"])

	// It's safe to manipulate map elements that haven't been explicitly assigned yet
	counters := make(map[string]int)
	counters["a"]++
	counters["a"]++
	counters["c"]++
	// counters["b"] hasn't been assigned but will display 0
	fmt.Println(counters["a"], counters["b"], counters["c"])

	// this won't work as the map hasn't been made
	// var nilMap map[int]string
	// fmt.Printf("%#v\n", nilMap)
	// nilMap[3] = "three"

	// var myMap map[int]string = make(map[int]string)
	myMap := make(map[int]string)
	myMap[3] = "three"
	fmt.Printf("%#v\n", myMap)

	checkDefaults()

	// 0 is assigned for Alma
	status("Alma")

	// Carl isn't assigned
	status("Carl")
}

func deleteItems() {
	var ok bool
	ranks := make(map[string]int)
	var rank int
	ranks["bronze"] = 3
	rank, ok = ranks["bronze"]
	fmt.Printf("rank: %d, ok: %v\n", rank, ok)

	// delete entry
	delete(ranks, "bronze")
	rank, ok = ranks["bronze"]
	fmt.Printf("rank: %d, ok: %v\n", rank, ok)

	isPrime := make(map[int]bool)
	var prime bool
	isPrime[5] = true
	prime, ok = isPrime[5]
	fmt.Printf("prime: %v, ok: %v\n", prime, ok)

	delete(isPrime, 5)
	prime, ok = isPrime[5]
	fmt.Printf("prime: %v, ok: %v\n", prime, ok)

}

func checkDefaults() {
	counters := map[string]int{"a": 3, "b": 0}
	var value int
	var ok bool
	// accessing an assigned value so ok is true
	value, ok = counters["a"]
	fmt.Println(value, ok)

	// accessed an assigned value so will be true
	value, ok = counters["b"]
	fmt.Println(value, ok)

	// Unassigned value, so ok false
	value, ok = counters["c"]
	fmt.Println(value, ok)

	// check whether value exists
	_, ok = counters["a"]
	fmt.Println(ok)

	// accessed an assigned value so will be true
	_, ok = counters["b"]
	fmt.Println(ok)

	// Unassigned value, so ok false
	_, ok = counters["c"]
	fmt.Println(ok)
}

func status(name string) {
	grades := map[string]float64{"Alma": 0, "Rohit": 86.5}
	grade, ok := grades[name]
	if !ok {
		fmt.Printf("No grade recorded for %s.\n", name)
	} else if grade < 60 {
		fmt.Printf("%s is failing!\n", name)
	}
}
