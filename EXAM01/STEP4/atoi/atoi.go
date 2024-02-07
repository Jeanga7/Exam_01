package main

import "fmt"

func Atoi(s string) int {
	var result int
	sign := 1

	for i, c := range s {
		if c == '+' && i ==0 {
			sign = 1
		}else	if c == '-' && i ==0 {
			sign = -1
		} else if c >= '0' && c <= '9' {
			result = result*10 + int(c-'0')
		} else {
			return 0
		}
	}

	return result * sign
}

func main() {
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi("--1234"))
}
