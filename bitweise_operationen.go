package main

import "fmt"

func main() {
	var a uint8 = 0b101010
	var b uint8 = 0b001101
	var c uint8 = 0b111111

	fmt.Printf("         a: %08b\n", a)
	fmt.Printf("       NOT: %08b\n", ^a)
	fmt.Println()
	fmt.Printf("a    AND b: %08b\n", a&b)
	fmt.Printf("a     OR b: %08b\n", a|b)
	fmt.Printf("a    XOR b: %08b\n", a^b)
	fmt.Printf("a ANDNOT b: %08b\n", a&^b)
	fmt.Printf("a ANDNOT c: %08b\n", a&^c)
	fmt.Printf("c ANDNOT a: %08b\n", c&^a)
}
