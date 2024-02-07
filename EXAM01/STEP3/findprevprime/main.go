package main

import "fmt"

func FindPrevPrime(nb int) int {
	res := 0
	for i := 0; i <= nb; i++ {
		if IsPrime(i){
			res = i
		}
	}
	return res
}

func IsPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}


func main() {
	fmt.Println(FindPrevPrime(5))
	fmt.Println(FindPrevPrime(4))
}
