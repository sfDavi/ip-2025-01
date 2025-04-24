package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	var MENOR, INTER, MAIOR int

	fmt.Print("Digite o valor de A: ")
	fmt.Scanln(&a)

	fmt.Print("Digite o valor de B: ")
	fmt.Scanln(&b)

	fmt.Print("Digite o valor de C: ")
	fmt.Scanln(&c)

	if a < b && a < c {
		MENOR = a
	} else if b < a && b < c {
		MENOR = b
	} else {
		MENOR = c
	}

	if a > b && a > c {
		MAIOR = a
	} else if b > a && b > c {
		MAIOR = b
	} else {
		MAIOR = c
	}

	if a > MENOR && a < MAIOR {
		INTER = a
	} else if b > MENOR && b < MAIOR {
		INTER = b
	} else {
		INTER = c
	}

	fmt.Printf("%d, %d, %d\n", MENOR, INTER, MAIOR)
}
