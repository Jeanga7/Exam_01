package main

import "fmt"

func Chunk(a []int, ch int) {
	if ch <= 0 {
		fmt.Println()
		return
	}
	result := make([][]int, 0, len(a)/ch)
	var slice []int
	for len(a) >= ch {
		slice, a = a[:ch], a[ch:]
		result = append(result, slice)
	}
	result = append(result, a[:])
	fmt.Println(result)
}

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}

// ====================================================
/* package main

import (
	"github.com/01-edu/z01"
)

func Chunk(a []int, ch int) {
	if ch <= 0 {
		z01.PrintRune('\n')
		return
	}

	result := make([][]int, 0, len(a)/ch)
	for len(a) >= ch {
		result, a = append(result, a[:ch]), a[ch:]
	}
	if len(a) > 0 {
		result = append(result, a[:])
	}

	for _, slice := range result {
		z01.PrintRune('[')
		for i, val := range slice {
			if i > 0 {
				z01.PrintRune(' ')
			}
			z01.PrintRune(rune(val + '0')) // Afficher chaque élément de la sous-slice.
		}
		z01.PrintRune(']')
	}

	z01.PrintRune('\n')
}

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}
*/
