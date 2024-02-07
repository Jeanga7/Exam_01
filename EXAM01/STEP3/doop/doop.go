package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	arg1, err1 := strconv.Atoi(os.Args[1])
	operator := os.Args[2]
	arg2, err2 := strconv.Atoi(os.Args[3])

	if err1 != nil || err2 != nil {
		return
	}

	var result int
	var err string

	if operator == "+" {
		result = arg1 + arg2
	} else if operator == "-" {
		result = arg1 - arg2
	} else if operator == "*" {
		result = arg1 * arg2
	} else if operator == "/" {
		if arg2 == 0 {
			err = "No division by 0"
			Print(err)
			return
		}
		result = arg1 / arg2
	} else if operator == "%" {
		if arg2 == 0 {
			err = "No modulo by 0"
			Print(err)
			return
		}
		result = arg1 % arg2
	} else {
		return
	}

	printNumber(result)
	z01.PrintRune('\n')
}

func Print(str string) {
	for _, v := range str {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}

func printNumber(n int) {
	result := Itoa(n)
	Print(result)
}

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}

	isNegative := false
	if n < 0 {
		isNegative = true
		n = -n
	}
	var result string
	for n != 0 {
		digit := n % 10
		result = string('0'+digit) + result
		n /= 10
	}

	if isNegative {
		z01.PrintRune('-')
	}

	return result
}

