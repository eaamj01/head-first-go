package main

import "fmt"

func main() {
	repeatLine("hello", 5)
	var width, height, area float64
	width = 4.2
	height = 3.0
	area = width * height
	fmt.Println(area/10.0, "litres needed")
	width = 5.2
	height = 3.5
	area = width * height
	fmt.Println(area/10.0, "litres needed")

	fmt.Println("About one-third:", 1.0/3.0)
	// format to 2 decimal places
	fmt.Printf("About one-third: %0.2f\n", 1.0/3.0)

	// Set to string
	resultString := fmt.Sprintf("About one-third: %0.2f\n", 1.0/3.0)
	fmt.Printf(resultString)

	// string and decimal int
	fmt.Printf("The %s cost %d pence each.\n", "gumballs", 23)

	// Insert float
	fmt.Printf("That will be $%f please\n", 0.23*5)

	// %f	Floating-point number
	// %d	Decimal integer
	// %s	String
	// %t	Boolean (true or false)
	// %v	Any value (chooses an appropriate format based on the supplied value’s type)
	// %#v	Any value, formatted as it would appear in Go program code
	// %T	Type of the supplied value (int, string, etc.)
	// %%	A literal percent sign
	fmt.Printf("A float: %f\n", 3.145)
	fmt.Printf("An integer: %d\n", 15)
	fmt.Printf("A string: %s\n", "hello")
	fmt.Printf("A boolean: %t\n", true)
	fmt.Printf("Values: %v %v %v \n", 1.2, "\t", true)
	fmt.Printf("Values %#v %#v %#v \n", 1.2, "\t", true)
	fmt.Printf("Values %#v %#v %#v \n", "", "\t", "\n")
	fmt.Printf("Types: %T %T %T\n", 1.2, "\t", true)
	fmt.Printf("Percent sign: %%\n")

	// minimum width of 12 and 2
	fmt.Printf("%12s | %2d\n", "Stamps", 50)
	fmt.Printf("%12s | %2d\n", "Paper Clips", 5)
	fmt.Printf("%12s | %2d\n", "Tape", 99)

	// format decimal places - %5 minimum width of whole number, 3f width after decimal point
	fmt.Printf("%5.3f\n", 12.655675) // 12.656

	// rounded to 3 places
	fmt.Printf("%%7.3f: %7.3f\n", 12.3456)

	// rounded to 2 places
	fmt.Printf("%%7.2f: %7.2f\n", 12.3456)

	// rounded to 1 place
	fmt.Printf("%%7.1f: %7.1f\n", 12.3456)

	// rounded to 1 place no padding
	fmt.Printf("%%.2ff: %.1f\n", 12.3456)

	// rounded to 1 place no padding
	fmt.Printf("%%.2f: %.2f\n", 12.3456)

	// print 1.82
	fmt.Printf("%.2f\n", 1.819999999999999999999999999998)
}

func repeatLine(line string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(line)
	}
}
