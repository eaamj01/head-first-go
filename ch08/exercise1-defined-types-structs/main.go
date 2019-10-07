package main

import "fmt"

type rectangle struct {
	length float64
	width  float64
}

func rectangleInfo(r rectangle) {
	fmt.Println("Length:", r.length)
	fmt.Println("Width:", r.width)
}

func makeSquare(r *rectangle) {
	if r.length > r.width {
		r.length = r.width
	} else {
		r.width = r.length
	}
}

func main() {
	rectangle := rectangle{length: 100, width: 50}
	rectangleInfo(rectangle)
	makeSquare(&rectangle)
	rectangleInfo(rectangle)
	rectangle.width = 100
	rectangle.length = 75
	makeSquare(&rectangle)
	rectangleInfo(rectangle)

}
