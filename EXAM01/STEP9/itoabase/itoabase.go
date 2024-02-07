package main

import (
	"fmt"
)

func ItoaBase(value, base int) string {
	if base < 2 || base > 16 {
		return ""
	}

	isNegative := false
	if value < 0 {
		isNegative = true
		value = -value
	}

	digits := "0123456789ABCDEF"
	result := ""
	for value > 0 {
		remainder := value % base
		result = string(digits[remainder]) + result
		value /= base
	}

	if isNegative {
		result = "-" + result
	}

	if result == "" {
		return "0"
	}

	return result
}

func main() {
	fmt.Println(ItoaBase(42, 2))   // Output: 101010
	fmt.Println(ItoaBase(-42, 16)) // Output: -2A
	fmt.Println(ItoaBase(255, 16)) // Output: FF
	fmt.Println(ItoaBase(123, 8))  // Output: 173
	fmt.Println(ItoaBase(456, 10)) // Output: 456
	fmt.Println(ItoaBase(987, 16)) // Output: 3DB
	fmt.Println(ItoaBase(0, 5))    // Output: 0
	fmt.Println(ItoaBase(123, 1))  // Output: ""
	fmt.Println(ItoaBase(123, 17)) // Output: ""
}

/* 
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
} */