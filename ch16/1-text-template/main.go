package main

import (
	"log"
	"os"
	"text/template"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// template text
	text := "Here's my template!\n"
	// create new template based on text
	tmpl, err := template.New("test").Parse(text)
	check(err)
	// write out the template text, Instead of an HTTP response, write to terminal
	err = tmpl.Execute(os.Stdout, nil)
	check(err)
}
