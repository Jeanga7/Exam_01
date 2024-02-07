package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		GetLastWord(os.Args[1])
	}
}

func GetLastWord(str string) {
	var word string
	for i, v := range str {
		if v == ' ' {
			for _, ch := range str[i:] {
				if ch != ' ' {
					word = str[i+1:]
				}
			}
		}
	}

	res := ""
	for _, v := range word {
		if v != ' ' {
			res += string(v)
		}
	}
	if len(res)>0{
		println(res)
	}
}
func println(str string){
	for _, v := range str {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
/*
func main() {
	if len(os.Args) != 2 {
		return
	}

	lastWord := getLastWord(os.Args[1])
	Print(lastWord)
}

func getLastWord(s string) string {
	var words []string
	var currentWord string
	inWord := false

	for _, char := range s {
		if char == ' ' || char == '\t' || char == '\n' {
			if inWord {
				words = append(words, currentWord)
				currentWord = ""
			}
			inWord = false
		} else {
			currentWord += string(char)
			inWord = true
		}
	}

	if inWord {
		words = append(words, currentWord)
	}

	if len(words) == 0 {
		return ""
	}

	return words[len(words)-1]
}

func Print(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}

*/
