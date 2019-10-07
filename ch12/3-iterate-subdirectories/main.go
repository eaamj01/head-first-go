package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func reportPanic() {
	p := recover()
	// if recover returned nil there is no panic
	if p == nil {
		// so do nothing
		return
	}
	// otherwise get the underlyingg error value and print it
	err, ok := p.(error)
	if ok {
		fmt.Println(err)
	} else {
		// if the panic values isn't an error, resume panicking with the same value
		panic(p)
	}
}
func scanDirectory(path string) {
	// print current directory
	fmt.Println(path)
	// get a slice with directory's contents
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		// join the directory path and filename
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			scanDirectory(filePath)
		} else {
			fmt.Println(filePath)
		}
	}
}

func main() {
	defer reportPanic()
	scanDirectory("/Users/andyjackson/go/src/github.com/eaamj01/head-first-go")

}
