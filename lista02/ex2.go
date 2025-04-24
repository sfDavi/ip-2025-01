package main

import (
	"fmt"
)

func main() {
	var numero int

	fmt.Print("Digite um número inteiro: ")
	fmt.Scanln(&numero)

	if numero > 0 {
		fmt.Println("O número inserido é POSITIVO.")
	} else if numero < 0 {
		fmt.Println("O número inserido é NEGATIVO.")
	} else {
		fmt.Println("O número inserido é NULO.")
	}
}
