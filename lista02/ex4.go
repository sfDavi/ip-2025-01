package main

import (
	"fmt"
	"math"
)

func main() {
	var numero float64

	fmt.Print("Digite um número: ")
	fmt.Scanln(&numero)

	if numero >= 0 {
		raiz := math.Sqrt(numero)
		fmt.Printf("A raiz quadrada do número é: %.2f\n", raiz)
	} else {
		quadrado := math.Pow(numero, 2)
		fmt.Printf("O quadrado do número é: %.2f\n", quadrado)
	}
}