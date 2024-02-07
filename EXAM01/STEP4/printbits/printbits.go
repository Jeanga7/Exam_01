package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func printBits(n int) {
	for i := 7; i >= 0; i-- {
		z01.PrintRune('0' + rune((n>>uint(i))&1))
	}
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		printBits(0)
		return
	}

	num, err := strconv.Atoi(args[0])
	if err != nil {
		printBits(0)
		return
	}

	printBits(num)
}