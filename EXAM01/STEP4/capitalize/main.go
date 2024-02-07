package main

import "fmt"

func Capitalize(s string) string {
	result := ""
	capitalizeNext := true

	for _, char := range s {
		if isAlphanumeric(char) {
			if capitalizeNext {
				result += string(ToUpper(char))
				capitalizeNext = false
			} else {
				result += string(ToLower(char))
			}
		} else {
			result += string(char)
			capitalizeNext = true
		}
	}

	return result
}

func isAlphanumeric(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')
}

func ToUpper(char rune) rune {
	if char >= 'a' && char <= 'z' {
		return char - ('a' - 'A')
	}
	return char
}

func ToLower(char rune) rune {
	if char >= 'A' && char <= 'Z' {
		return char + ('a' - 'A')
	}
	return char
}

func main() {
	fmt.Println(Capitalize("Hello! How are you? How+are+things+4you?"))
}
