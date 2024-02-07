package main

import "github.com/01-edu/z01"

func Rot14(s string) string {
	result := ""
	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			if v >= 'm' {
				v -= 12
			} else {
				v += 14
			}
		}else if v >= 'A' && v <= 'Z' {
			if v >= 'M' {
				v -= 12
			} else {
				v += 14
			}
		}
		result += string(v)
	}
	return result
}

func main() {
	result := Rot14("Hello! How are You?")

	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
