package main

import (
	"strconv"

	"github.com/01-edu/z01"
)

func ForEach(f func(int), a []int) {
	for _, value := range a {
		f(value)
	}
}

func PrintNbr(n int) {
	for _, digit := range strconv.Itoa(n) {
		z01.PrintRune(rune(digit))
	}
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	ForEach(PrintNbr, a)
	z01.PrintRune('\n')
}
