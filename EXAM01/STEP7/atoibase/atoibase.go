package main

import (
	"fmt"
)

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}

	seen := make(map[rune]struct{})
	for _, char := range base {
		if char == '+' || char == '-' {
			return false
		}
		if _, exists := seen[char]; exists {
			return false
		}
		seen[char] = struct{}{}
	}

	return true
}

func AtoiBase(s string, base string) int {
	if !isValidBase(base) {
		return 0
	}

	baseMap := make(map[rune]int)
	for i, char := range base {
		baseMap[char] = i
	}

	result := 0
	baseLen := len(base)
	for _, char := range s {
		if index, exists := baseMap[char]; exists {
			result = result*baseLen + index
		} else {
			return 0
		}
	}

	return result
}

func main() {
	fmt.Println(AtoiBase("125", "0123456789"))
	fmt.Println(AtoiBase("1111101", "01"))
	fmt.Println(AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(AtoiBase("uoi", "choumi"))
	fmt.Println(AtoiBase("bbbbbab", "-ab"))
}
