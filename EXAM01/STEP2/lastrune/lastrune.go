package main

import "github.com/01-edu/z01"

/* func LastRune(s string) rune {
	runes := []rune(s)
	if len(runes) > 0 {
		return runes[len(runes)-1]
	}
	return 0 
} */

func LastRune(s string) rune {
	return rune(s[len(s)-1])
}
func main() {
	z01.PrintRune(LastRune("Hello!"))
	z01.PrintRune(LastRune("Salut!"))
	z01.PrintRune(LastRune("Ola!"))
	z01.PrintRune('\n')
}
