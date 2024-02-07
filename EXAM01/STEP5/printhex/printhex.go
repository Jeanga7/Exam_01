package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Atoi(s string) int {
	var result int
	for _, c := range s {
	 if c >= '0' && c <= '9' {
			result = result*10 + int(c-'0')
		} else {
			return 0
		}
	}
	return result 
}

func printHex(number int) {
	const hexDigits = "0123456789abcdef"
	var result []rune

	if number == 0 {
		z01.PrintRune('0')
		return
	}

	for number > 0 {
		remainder := number % 16
		result = append(result, rune(hexDigits[remainder]))
		number /= 16
	}

	for i := len(result) - 1; i >= 0; i-- {
		z01.PrintRune(result[i])
	}
}


func main() {
	if len(os.Args) != 2 {
		return
	}

	result := Atoi(os.Args[1])
	if result == 0 {
		z01.PrintRune('E')
		z01.PrintRune('R')
		z01.PrintRune('R')
		z01.PrintRune('O')
		z01.PrintRune('R')
		z01.PrintRune('\n')
		return
	}
	printHex(result)
	z01.PrintRune('\n')
}
