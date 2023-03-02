package main

import (
	"fmt"
	"os"
)

func main() {
	// print decimal and the binary notation
	fmt.Printf("%3d: %08b\n", 0, 0)
	fmt.Printf("%3d: %08b\n", 1, 1)
	fmt.Printf("%3d: %08b\n", 2, 2)
	fmt.Printf("%3d: %08b\n", 3, 3)
	fmt.Printf("%3d: %08b\n", 4, 4)
	fmt.Printf("%3d: %08b\n", 5, 5)
	fmt.Printf("%3d: %08b\n", 6, 6)
	fmt.Printf("%3d: %08b\n", 7, 7)
	fmt.Printf("%3d: %08b\n", 8, 8)
	fmt.Printf("%3d: %08b\n", 16, 16)
	fmt.Printf("%3d: %08b\n", 32, 32)
	fmt.Printf("%3d: %08b\n", 64, 64)
	fmt.Printf("%3d: %08b\n", 128, 128)

	// bitwise and operator
	fmt.Println("bitwise and operator")
	fmt.Printf("%b & %b == %b\n", 0, 0, 0&0)
	fmt.Printf("%b & %b == %b\n", 0, 1, 0&1)
	fmt.Printf("%b & %b == %b\n", 1, 1, 1&1)

	fmt.Printf("%02b\n", 1)
	fmt.Printf("%02b\n", 3)
	fmt.Printf("%02b\n", 1&3)

	fmt.Printf("%02b\n", 2)
	fmt.Printf("%02b\n", 3)
	fmt.Printf("%02b\n", 2&3)

	fmt.Printf("%08b\n", 170)
	fmt.Printf("%08b\n", 15)
	fmt.Printf("%08b\n", 170&15)

	// bitwise OR operator
	fmt.Println("bitwise OR operator")
	fmt.Printf("%b | %b == %b\n", 0, 0, 0|0)
	fmt.Printf("%b | %b == %b\n", 0, 1, 0|1)
	fmt.Printf("%b | %b == %b\n", 1, 1, 1|1)

	fmt.Printf("%02b\n", 1)
	fmt.Printf("%02b\n", 0)
	fmt.Printf("%02b\n", 1|0)

	fmt.Printf("%02b\n", 2)
	fmt.Printf("%02b\n", 0)
	fmt.Printf("%02b\n", 2|0)

	fmt.Printf("%08b\n", 170)
	fmt.Printf("%08b\n", 15)
	fmt.Printf("%08b\n", 170|15)

	fmt.Printf(" os.O_RDONLY %016b\n", os.O_RDONLY)
	fmt.Printf(" os.O_WRONLY %016b\n", os.O_WRONLY)
	fmt.Printf(" os.O_RDWR %016b\n", os.O_RDWR)
	fmt.Printf(" os.O_CREATE %016b\n", os.O_CREATE)
	fmt.Printf(" os.O_APPEND %016b\n", os.O_APPEND)
	fmt.Printf(" os.O_WRONLY|os.O_CREATE %016b\n", os.O_WRONLY|os.O_CREATE)
	fmt.Printf(" os.O_WRONLY|os.O_CREATE||os.O_APPEND %016b\n", os.O_WRONLY|os.O_CREATE|os.O_APPEND)
}
