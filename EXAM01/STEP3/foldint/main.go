package main

import (
	"fmt"
	/* "strconv"

	"github.com/01-edu/z01" */)

func FoldInt(f func(int, int) int, a []int, n int) {
	for _, value := range a {
		n = f(n, value)
	}
	fmt.Println(n)
}

/* func FoldInt(f func(int, int) int, a []int, n int) {
	for _, value := range a {
		n = f(n, value)
	}

	resultStr := strconv.Itoa(n)

	for _, r := range resultStr {
		z01.PrintRune(r)
	}

	z01.PrintRune('\n')
} */

func Add(a, b int) int {
	return a + b
}

func Mul(a, b int) int {
	return a * b
}

func Sub(a, b int) int {
	return a - b
}

func main() {
	table := []int{1, 2, 3}
	ac := 93
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
	fmt.Println()

	table = []int{0}
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
}
