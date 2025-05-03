package main

import (
	"fmt"
)

func main() {
	var x int
	var xb int
	fmt.Print("Digite um número: ")
	fmt.Scan(&x)
	xb = convertebinario(x)
	fmt.Printf("Forma binária: %d", xb)
}

func convertebinario(x int) int {
	if x == 0 {
		return 0
	}
	return x%2 + 10*convertebinario(x/2)
}
