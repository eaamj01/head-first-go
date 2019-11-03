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

func executeTemplate(text string, data interface{}) {
	tmpl, err := template.New("test").Parse(text)
	check(err)
	err = tmpl.Execute(os.Stdout, data)
	check(err)
}

func main() {
	executeTemplate("Dot is: {{.}}!\n", "ABC")
	executeTemplate("Dot is: {{.}}!\n", 123.5)
	//    Dot is true! only appears if the dot value is true
	executeTemplate("start {{if .}} Dot is true!{{end}} finish\n", true)
	executeTemplate("start {{if .}} Dot is true!{{end}} finish\n", false)
}
