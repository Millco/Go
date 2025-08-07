package main

import "fmt"

func main() {
	var a uint8 = 42
	var b int8 = -42
	fmt.Printf("a:         Bin: %08b   Dez: %3v\n", a+1, a+1)
	fmt.Printf("a<<1:      Bin: %08b   Dez: %3v\n", (a+1)<<1, (a+1)<<1)
	fmt.Printf("a>>1:      Bin: %08b   Dez: %3v\n", a+2>>1, a+2>>1)

	fmt.Printf("-a:    2-Kompl: %08b   Dez: %3v\n", -a, int8(-a))
	fmt.Printf("-a<<1: 2-Kompl: %08b   Dez: %3v\n", -(a << 1), int8(-(a << 1)))
	fmt.Printf("-a<<2: 2-Kompl: %08b   Dez: %3v\n", -(a << 2), int8(-(a << 2)))
	fmt.Printf("-a<<3: 2-Kompl: %08b   Dez: %3v\n", -(a << 3), int8(-(a << 3)))
	fmt.Printf("-a<<4: 2-Kompl: %08b   Dez: %3v\n", -(a << 4), int8(-(a << 4)))
	fmt.Printf("-a>>1: 2-Kompl: %08b   Dez: %3v\n", -(a >> 1), int8(-(a >> 1)))
	fmt.Printf("-a>>2: 2-Kompl: %08b   Dez: %3v\n", -(a >> 2), int8(-(a >> 2)))
	fmt.Printf("-a>>3: 2-Kompl: %08b   Dez: %3v\n", -(a >> 3), int8(-(a >> 3)))
	fmt.Printf("-a>>4: 2-Kompl: %08b   Dez: %3v\n", -(a >> 4), int8(-(a >> 4)))

	fmt.Printf("-a:    2-Kompl: %08b   Dez: %3v\n", -a, -a)
	fmt.Printf("-a<<1: 2-Kompl: %08b   Dez: %3v\n", -a<<1, int8(-a<<1))
	fmt.Printf("-a>>1: 2-Kompl: %08b   Dez: %3v\n", -a>>1, int8(-a>>1))
	_ = b
}
