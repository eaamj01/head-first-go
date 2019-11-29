package main

import (
	"fmt"
	"os"
)

func main() {

	// print decimal and octal version
	for i := 0; i <= 19; i++ {
		fmt.Printf("%3d: %04o\n", i, i)
	}

	// Prefixing an int with 0 writes the number in octal notation
	fmt.Printf("Decimal 1: %3d Octal 01:%2d\n", 1, 01)
	fmt.Printf("Decimal 10: %3d Octal 010:%2d\n", 10, 010)
	fmt.Printf("Decimal 100: %3d Octal 0100:%2d\n", 100, 0100)

	// converting octal to file permissions
	fmt.Printf("%09b\n", 0007)
	fmt.Printf("%09b\n", 0070)
	fmt.Printf("%09b\n", 0700)

	//          | print the binary representation of an octal number
	//          |     | print the string for the FileMode conversion of the same number
	//          \/    \/
	fmt.Printf("%09b %s\n", 0000, os.FileMode(0000))
	fmt.Printf("%09b %s\n", 0111, os.FileMode(0111))
	fmt.Printf("%09b %s\n", 0222, os.FileMode(0222))
	fmt.Printf("%09b %s\n", 0333, os.FileMode(0333))
	fmt.Printf("%09b %s\n", 0444, os.FileMode(0444))
	fmt.Printf("%09b %s\n", 0555, os.FileMode(0555))
	fmt.Printf("%09b %s\n", 0666, os.FileMode(0666))
	fmt.Printf("%09b %s\n", 0777, os.FileMode(0777))
}
