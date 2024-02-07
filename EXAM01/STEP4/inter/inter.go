package main

import (
	"os"

	"github.com/01-edu/z01"
)

func inter(s1, s2 string) {
	var order []rune
	mapChar := make(map[rune]bool)
	for _, char := range s1 {
		if exist := mapChar[char]; !exist {
			order = append(order, char)
			mapChar[char] = true
		}
	}

	var result []rune
	for _, char := range s2 {
		if mapChar[char]{
			result = append(result, char)
		}
	}

	for _, char := range order {
		if !contains(result, char){
			z01.PrintRune(char)
		}
	}
	z01.PrintRune('\n')
}

func contains(res []rune, target rune) bool{
	for _, char := range res {
		if char == target {
			return false
		}
	}
	return true
}


func main() {
	if len(os.Args) == 3 {
		inter(os.Args[1], os.Args[2])
	}
}
