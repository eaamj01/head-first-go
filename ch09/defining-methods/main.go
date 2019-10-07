package main

import "fmt"

// MyType A Receiver Parameter
type MyType string

func main() {
	example1 := MyType("A MyType value")
	example1.sayHi()

	anotherValue := MyType("Another MyType value")
	anotherValue.sayHi()

	example1.MethodWithParameters(241, true)
	fmt.Println("Length of m:", example1.MethodWithReturn())

}

// MyType is the receiver parameter and sayHi() the method defoned on MyType Value
func (m MyType) sayHi() {
	fmt.Println("Hi from, m")
}

// MethodWithParameters Example of MyType method that has parameters
func (m MyType) MethodWithParameters(number int, flag bool) {
	fmt.Println(m)
	fmt.Println(number)
	fmt.Println(flag)
}

// MethodWithReturn Example of MyType method that has a return value
func (m MyType) MethodWithReturn() int {
	return len(m)
}
