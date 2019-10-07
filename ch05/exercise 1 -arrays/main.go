package main

import "fmt"

func main() {
	arrWeekdays := [5]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}
	for _, weekDay := range arrWeekdays {
		fmt.Println(weekDay)
	}
}
