package main

import (
	"fmt"
)

func SwapBits(octet byte) byte {
	return (octet>>4) | (octet<<4)
}
/* func SwapBits(octet byte) byte {
	// Déplacer les 4 bits de gauche vers la droite
	left := octet >> 4

	// Déplacer les 4 bits de droite vers la gauche
	right := octet << 4

	// Utiliser l'opérateur OR pour combiner les deux moitiés
	result := left | right

	return result
} */

func main() {
	octet := byte(0x41) // La valeur hexadécimale 0x41 est 01000001 en binaire
	result := SwapBits(octet)
	fmt.Printf("%08b | %08b\n", octet, result)
}