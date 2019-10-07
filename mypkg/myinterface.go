package mypkg

import "fmt"

type MyInterface interface {
	MethodWithoutParameters()
	MethodWithParameter(float64)
	MethodWithReturnValue() string
}

// MyType Will make this satisfy the interface
type MyType int

// MethodWithoutParameters First required method
func (m MyType) MethodWithoutParameters() {
	fmt.Println("MethodWithoutParameters called")
}

// MethodWithParameter Second required method
func (m MyType) MethodWithParameter(f float64) {
	fmt.Println("MethodWithParameter called with", f)
}

// MethodWithReturnValue Third required method (with string return)
func (m MyType) MethodWithReturnValue() string {
	return "Hi from MethodWithReturnValue"
}

// MethodNotInInterface A type can still satisfy an interface even
// if it has methods that aren't part of the interface
func (m MyType) MethodNotInInterface() {
	fmt.Println("MethodNotInInterface called")
}
