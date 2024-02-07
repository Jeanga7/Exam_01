package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isSpace(char rune) bool {
	return char == ' ' || char == '\t' || char == '\n' || char == '\r'
}

func cleanStr(s string) string {
	// Supprimer les espaces et les tabulations au début et à la fin de la chaîne
	start := 0
	end := len(s) - 1

	for start <= end && isSpace(rune(s[start])) {
		start++
	}

	for end >= start && isSpace(rune(s[end])) {
		end--
	}

	// Construire la nouvelle chaîne sans espaces indésirables
	result := ""
	for i := start; i <= end; i++ {
		result += string(s[i])
		if isSpace(rune(s[i])) {
			for i <= end && isSpace(rune(s[i])) {
				i++
			}
			i--
		}
	}

	return result
}

func main() {
	// Vérifier le nombre d'arguments
	if len(os.Args) != 2 {
		z01.PrintRune('\n')
		return
	}

	// Appeler la fonction cleanStr avec l'argument
	result := cleanStr(os.Args[1])

	// Afficher le résultat suivi d'une nouvelle ligne
	for _, char := range result {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
