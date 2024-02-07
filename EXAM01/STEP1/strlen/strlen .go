package main

import "fmt"

/* func StrLen(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		count++
	}
	return count
} */

func StrLen(s string) int {
	return len(s)
}


func main() {
	l := StrLen("Hello World!")
	fmt.Println(l)
}
