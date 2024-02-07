package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	result := Rot13(os.Args[1])
	Print(result)
}
func Rot13(s string) (result string) {
	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			if v >= 'm' {
				v -= 13
			} else {
				v += 13
			}
		} else if v >= 'A' && v <= 'Z' {
			if v >= 'M' {
				v -= 13
			} else {
				v += 13
			}
		}
		result += string(v)
	}
	return result
}


func Print(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
