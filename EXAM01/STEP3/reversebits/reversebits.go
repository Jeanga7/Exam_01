package main

import "fmt"

func ReverseBits(oct byte) byte {
	var result byte

	for i := uint(0); i < 8; i++ {
		result |= (oct >> i & 1) << (7 - i)
	}
	return result
}

func main() {
	oct := byte(38)
	fmt.Printf("Original: %08b\n", oct)
	fmt.Printf("Reversed: %08b\n", ReverseBits(oct))
}
