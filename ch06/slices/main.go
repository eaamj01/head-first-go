package main

import "fmt"

// a slice is a view to an underlying array
func main() {
	// declare a slice variable
	var notes []string

	// create a slice with 7 strings
	notes = make([]string, 7)
	fmt.Println(notes)

	// assign and create slice
	primes := make([]int, 5)
	primes[0] = 2
	primes[1] = 2
	fmt.Println(primes[0])

	fmt.Println(len(notes))
	fmt.Println(len(primes))

	sliLetters := []string{"a", "b", "c"}
	for i := 0; i < len(sliLetters); i++ {
		fmt.Println(sliLetters[i])
	}

	for _, letters := range sliLetters {
		fmt.Println(letters)
	}

	underLyingArray := [5]string{"a", "b", "c", "d", "e"}
	// the second element is the number to stop before
	slice1 := underLyingArray[0:3] // sets [a b c]
	fmt.Println(slice1)

	underLyingArray2 := [5]string{
		"a",
		"b", // get index 1 through
		"c",
		"d", // to the one before 4
		"e"}
	i, j := 1, 4
	slice2 := underLyingArray2[i:j]
	fmt.Println(slice2)

	underLyingArray3 := [5]string{
		"a", // get start index through
		"b",
		"c", // to the one before 3
		"d",
		"e"}

	// omit start, defaults to
	slice3 := underLyingArray3[:3]
	fmt.Println(slice3)

	//---------------------------

	underLyingArray4 := [5]string{
		"a", // get start index through
		"b",
		"c", // start here through to the nd
		"d",
		"e"}

	// omit end, gets all from start position
	slice4 := underLyingArray4[2:]
	fmt.Println(slice4)

	slice := []string{"a", "b"}
	fmt.Println(slice, len(slice))
	slice = append(slice, "c")
	fmt.Println(slice, len(slice))
	slice = append(slice, "d", "e") // append 2 elements
	fmt.Println(slice, len(slice))

	// slices should be assigned back to original slice as underlying array may not be able to expand
	// or end up with copies and inconsistancies
	s1 := []string{"s1", "s1"}
	s2 := append(s1, "s2", "s2")
	s3 := append(s2, "s3", "s3")
	s4 := append(s3, "s4", "s4")
	fmt.Println(s1, s2, s3)
	s4[0] = "XX"
	fmt.Println(s1, s2, s3, s4)
	// [s1 s1] [s1 s1 s2 s2] [s1 s1 s2 s2 s3 s3]
	//   s1         s2              s3                     s4
	// [s1 s1] [s1 s1 s2 s2] [XX s1 s2 s2 s3 s3] [XX s1 s2 s2 s3 s3 s4 s4]
	// s1 and s2 have different underlying arrays
	// s3 slice shares same underlying array as s4

	// initialise as int or false
	floatSlice := make([]float64, 10)
	boolSlice := make([]bool, 10)
	fmt.Println(floatSlice[9], boolSlice[5])

	// int and string assigned as nil
	var intSlice []int
	var stringSlice []string
	fmt.Printf("intSlice: %#v, stringSlice: %#v\n", intSlice, stringSlice)

	// if len is passed a nil slice, length is 0
	fmt.Println(len(intSlice))

	// if an item is appended to a nil slice, then that becomes a 1 item slice
	intSlice = append(intSlice, 27)
	fmt.Printf("intSlice: %#v\n", intSlice)

	var slString []string
	if len(slString) == 0 {
		slString = append(slString, "first item")
	}
	fmt.Printf("%#v\n", slString)
}
