package main

import (
	"fmt"
)

func main() {
	var numero1, numero2, soma int

	fmt.Print("Digite o primeiro número inteiro: ")
	fmt.Scanln(&numero1)

	fmt.Print("Digite o segundo número inteiro: ")
	fmt.Scanln(&numero2)

	soma = numero1 + numero2

	if soma > 20 {
		soma += 8
		fmt.Printf("O resultado é: %d\n", soma)
	} else {
		soma -= 5
		fmt.Printf("O resultado é: %d\n", soma)
	}
}
