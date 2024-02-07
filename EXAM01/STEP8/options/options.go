package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printOptions(options int) {
	for i := 25; i >= 0; i-- {
		bit := (options >> uint(i)) & 1
		z01.PrintRune(rune('0' + bit))
		if i%8 == 0 {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}

func isValidOption(arg string) bool {
	if arg == "-h" {
		printAllOptions()
		return false
	}
	if arg[0] == '-' && len(arg) > 1 {
		validChars := "abcdefghijklmnopqrstuvwxyz"
		for _, char := range arg[1:] {
			if !contains(validChars, char) {
				return false
			}
		}
		return true
	}
	return false
}

func contains(s string, char rune) bool {
	for _, c := range s {
		if c == char {
			return true
		}
	}
	return false
}

func printAllOptions() {
	optionsString := "abcdefghijklmnopqrstuvwxyz"
	for i := len(optionsString) - 1; i >= 0; i-- {
		z01.PrintRune(rune(optionsString[i]))
	}
	z01.PrintRune('\n')
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		printAllOptions()
		return
	}

	options := 0
	for _, arg := range args {
		if !isValidOption(arg) {
			printInvalidOption()
			return
		}
		if arg == "-h" {
			return
		}
		options = 0 // Réinitialiser les options pour chaque nouvelle option spécifiée
		for _, char := range arg[1:] {
			options |= 1 << uint(char-'a')
		}
	}

	printOptions(options)
}

func printInvalidOption() {
	for _, char := range "Invalid Option\n" {
		z01.PrintRune(char)
	}
}
