package main

import (
	"fmt"

	"github.com/eaamj01/head-first-go/magazine"
)

func main() {
	subscriber2 := magazine.Subscriber{Name: "Aman Singh"}
	subscriber2.HomeAddress.Street = "123 St"
	subscriber2.HomeAddress.City = "Omaha"
	subscriber2.HomeAddress.State = "NE"
	subscriber2.HomeAddress.PostalCode = "68111"

	fmt.Println("Subscriber Name:", subscriber2.Name)
	fmt.Println("Street:", subscriber2.HomeAddress.Street)
	fmt.Println("City:", subscriber2.HomeAddress.City)
	fmt.Println("State:", subscriber2.HomeAddress.State)
	fmt.Println("PostalCode:", subscriber2.HomeAddress.PostalCode)

	// Using Address  but this isn't necessay
	employee1 := magazine.Employee{Name: "Joy Carr"}
	employee1.Address.Street = "456 Elm St"
	employee1.Address.City = "Portland"
	employee1.Address.State = "OR"
	employee1.Address.PostalCode = "97222"

	fmt.Println("Subscriber Name:", employee1.Name)
	fmt.Println("Street:", employee1.Address.Street)
	fmt.Println("City:", employee1.Address.City)
	fmt.Println("State:", employee1.Address.State)
	fmt.Println("PostalCode:", employee1.Address.PostalCode)

	employee2 := magazine.Employee{Name: "Joy Carr"}
	employee2.Street = "456 Elm St"
	employee2.City = "Portland"
	employee2.State = "OR"
	employee2.PostalCode = "97222"

	fmt.Println("Subscriber Name:", employee2.Name)
	fmt.Println("Street:", employee2.Street)
	fmt.Println("City:", employee2.City)
	fmt.Println("State:", employee2.State)
	fmt.Println("PostalCode:", employee2.PostalCode)

}
