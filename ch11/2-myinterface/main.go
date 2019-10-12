package main

import "github.com/eaamj01/head-first-go/mypkg"

func main() {
	var value mypkg.MyInterface

	value = mypkg.MyType(5)
	value.MethodWithoutParameters()
	value.MethodWithParameter(127.3)
	value.MethodWithReturnValue()
}
