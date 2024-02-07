package main

import "fmt"

func Max(a []int) (max int) {
	for _, v := range a {
		if v > max{
			max = v
		}
	}
	return max
}

func main() {
	a := []int{23, 123, 1, 11, 55, 93}
	max := Max(a)
	fmt.Println(max)
}
