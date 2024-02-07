package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isOpeningBracket(char rune) bool {
	return char == '(' || char == '[' || char == '{'
}

func isClosingBracket(char rune) bool {
	return char == ')' || char == ']' || char == '}'
}

func isMatchingPair(opening rune, closing rune) bool {
	switch opening {
	case '(':
		return closing == ')'
	case '[':
		return closing == ']'
	case '{':
		return closing == '}'
	default:
		return false
	}
}

func checkBrackets(expr string) bool {
	var stack []rune

	for _, char := range expr {
		if isOpeningBracket(char) {
			stack = append(stack, char)
		} else if isClosingBracket(char) {
			if len(stack) == 0 || !isMatchingPair(stack[len(stack)-1], char) {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func main() {
	args := os.Args[1:]

	for _, arg := range args {
		if checkBrackets(arg) {
			z01.PrintRune('O')
			z01.PrintRune('K')
		} else {
			z01.PrintRune('E')
			z01.PrintRune('r')
			z01.PrintRune('r')
			z01.PrintRune('o')
			z01.PrintRune('r')
		}
		z01.PrintRune('\n')
	}
}
