package main

import (
	"fmt"
	"log"

	"github.com/eaamj01/head-for-go/calendar"
)

func dateSetting() {
	date := calendar.Date{}
	err := date.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetMonth(7)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetDay(25)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(date.Day(), date.Month(), date.Year())
}

func main() {
	dateSetting()

	event := calendar.Event{}
	err := event.SetTitle("A Birthday")
	//err := event.SetTitle("An extremely long and impractical title")
	if err != nil {
		log.Fatal(err)
	}

	err = event.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}

	err = event.SetMonth(7)
	if err != nil {
		log.Fatal(err)
	}

	err = event.SetDay(25)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(event.Day(), event.Month(), event.Year())

}
