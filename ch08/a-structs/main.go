package main

import "fmt"

type part struct {
	description string
	count       int
}

type car struct {
	name     string
	topSpeed float64
}

func main() {
	printMyStruct()
	subscriberAsVar()
	subscriberAsType()
	printCarStruct()
}

// printMyStruct myStruct is assigned as a variable and values are assigned to it
func printMyStruct() {
	var myStruct struct {
		number float64
		word   string
		toggle bool
	}

	fmt.Printf("%#v\n", myStruct)
	myStruct.number = 3.14
	myStruct.word = "pie"
	myStruct.toggle = true

	fmt.Println("Number:", myStruct.number)
	fmt.Println("Word:", myStruct.word)
	fmt.Println("Toggle:", myStruct.toggle)
	fmt.Printf("%#v\n", myStruct)
}

// subscriberAsVar An example of a struct that has only one use
func subscriberAsVar() {
	var subscriber struct {
		name   string
		rate   float64
		active bool
	}

	subscriber.name = "Aman Singh"
	subscriber.rate = 4.99
	subscriber.active = true
	fmt.Println("Name:", subscriber.name)
	fmt.Println("Rate:", subscriber.rate)
	fmt.Println("Active?:", subscriber.active)
}

// subscriberAsType Example of setting subscriber as a type which can then be assigned to variables
func subscriberAsType() {
	type subscriber struct {
		name   string
		rate   float64
		active bool
	}

	var subscriber1 subscriber
	subscriber1.name = "Aman Singh"
	subscriber1.rate = 4.99
	subscriber1.active = true

	fmt.Println("Name:", subscriber1.name)
	fmt.Println("Rate:", subscriber1.rate)
	fmt.Println("Active?:", subscriber1.active)

	var subscriber2 subscriber
	subscriber2.name = "Beth Ryan"
	subscriber2.rate = 2.99
	subscriber2.active = true

	fmt.Println("Name:", subscriber2.name)
	fmt.Println("Rate:", subscriber2.rate)
	fmt.Println("Active?:", subscriber2.active)
}

// printCarStruct Assign values to car and part global structs
func printCarStruct() {
	var porsche car
	porsche.name = "Porsche 911 R"
	porsche.topSpeed = 323

	fmt.Println("Name:", porsche.name)
	fmt.Println("Top speed:", porsche.topSpeed)

	var bolts part
	bolts.description = "Hex bolts"
	bolts.count = 24

	fmt.Println("Description:", bolts.description)
	fmt.Println("Count:", bolts.count)
}
