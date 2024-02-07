package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isVowel(char rune) bool {
	vowels := "aeiouAEIOU"
	for _, v := range vowels {
		if v == char {
			return true
		}
	}
	return false
}

func pigLatin(s string) string {
	var words []string
	var result string

	word := ""
	for _, char := range s {
		if char == ' ' {
			words = append(words, word)
			word = ""
		} else {
			word += string(char)
		}
	}
	words = append(words, word)

	for _, word := range words {
		if isVowel(rune(word[0])) {
			result += word + "ay "
		} else {
			firstVowelIndex := -1
			for i, char := range word {
				if isVowel(char) {
					firstVowelIndex = i
					break
				}
			}

			if firstVowelIndex == -1 {
				result += "No vowels "
			} else {
				result += word[firstVowelIndex:] + word[:firstVowelIndex] + "ay "
			}
		}
	}
	return result
}

func main() {
	if len(os.Args) != 2 {
		return
	}
	result := pigLatin(os.Args[1])

	for _, char := range result {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
