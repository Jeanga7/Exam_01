package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isAlphanumeric(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')
}

func reverseWords(s string) string {
	var reversedWords []string
	var currentWord []rune

	for _, char := range s {
		if isAlphanumeric(char) {
			currentWord = append(currentWord, char)
		} else if len(currentWord) > 0 {
			reversedWords = append(reversedWords, string(currentWord))
			currentWord = nil
		}
	}

	if len(currentWord) > 0 {
		reversedWords = append(reversedWords, string(currentWord))
	}

	// Inverser l'ordre des mots
	for i, j := 0, len(reversedWords)-1; i < j; i, j = i+1, j-1 {
		reversedWords[i], reversedWords[j] = reversedWords[j], reversedWords[i]
	}

	return concatenateWords(reversedWords)
}

func concatenateWords(words []string) string {
	var result string
	for i, word := range words {
		result += word
		if i < len(words)-1 {
			result += " "
		}
	}
	return result
}

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		return
	}

	reversed := reverseWords(args[0])

	for _, char := range reversed {
		z01.PrintRune(char)
	}

	z01.PrintRune('\n')
}
