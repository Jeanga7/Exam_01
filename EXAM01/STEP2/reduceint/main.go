package main

import (
	"github.com/01-edu/z01"
)

func ReduceInt(a []int, f func(int, int) int) {
	if len(a) == 0 {
		return
	}

	result := a[0]
	for i := 1; i < len(a); i++ {
		result = f(result, a[i])
	}

	for _, digit := range itoa(result) {
		z01.PrintRune(digit)
	}
	z01.PrintRune('\n')
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}

	var result string
	for n != 0 {
		digit := n % 10
		result = string('0'+digit) + result
		n /= 10
	}

	if n < 0 {
		result = "-" + result
	}

	return result
}


func main() {
	mul := func(acc int, cur int) int {
		return acc * cur
	}
	sum := func(acc int, cur int) int {
		return acc + cur
	}
	div := func(acc int, cur int) int {
		return acc / cur
	}

	as := []int{500, 2}
	ReduceInt(as, mul)
	ReduceInt(as, sum)
	ReduceInt(as, div)
}
