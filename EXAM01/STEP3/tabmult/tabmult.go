package main

import (
	"os"

	"github.com/01-edu/z01"
)

func tabmult(nbr int) {
	for i := 1; i < 10; i++ {
		result := nbr * i
		printTableLine(i, nbr, result)
	}
}
func printTableLine(i, nbr, result int) {
	print(itoa(i) + " x " + itoa(nbr) + " = " + itoa(result))
	z01.PrintRune('\n')
}

func main() {
	if len(os.Args) == 2 {
		number := Atoi(os.Args[1])
		tabmult(number)
	} else {
		z01.PrintRune('\n')
	}
}

func Atoi(s string) int {
	var result int
	sign := 1

	for i, c := range s {
		if c == '-' && i == 0 {
			sign = -1
		} else if c >= '0' && c <= '9' {
			result = result*10 + int(c-'0')
		} else {
			return 0
		}
	}

	return result * sign
}
func print(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
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
