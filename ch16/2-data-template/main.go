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
	templateText := "Template start \nAction: {{.}}\nTemplate end\n"
	tmpl, err := template.New("test").Parse(templateText)
	check(err)
	// write out the template text, Instead of an HTTP response, write to terminal
	err = tmpl.Execute(os.Stdout, "ABC")
	check(err)
	err = tmpl.Execute(os.Stdout, 42)
	check(err)
	err = tmpl.Execute(os.Stdout, true)
	check(err)
}
