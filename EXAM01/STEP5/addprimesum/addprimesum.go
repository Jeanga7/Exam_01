package main

import (
	"fmt"
	"os"
	"strconv"
)

func isPrime(num int) bool {
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

func addPrimeSum(n int) int {
	sum := 0
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			sum += i
		}
	}
	return sum
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("0")
		return
	}

	num, err := strconv.Atoi(os.Args[1])
	if err != nil || num <= 0 {
		fmt.Println("0")
		return
	}

	result := addPrimeSum(num)
	fmt.Println(result)
}
