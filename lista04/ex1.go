package main

import "fmt"

func main() {
	var numeros [10]int
	var num int
	fmt.Println("Digite 10 números inteiros:")

	for i := 0; i < 10; i++ {
		fmt.Scan(&numeros[i])
	}

	teste := false
	fmt.Println("Números superiores a 50 e suas posições:")
	for j := 0; j < 10; j++ {
		num = numeros[j]
		if num > 50 {
			fmt.Printf("Número: %d, Posição: %d\n", num, j)
			teste = true
		}
	}

	if !teste {
		fmt.Println("Nenhum número superior a 50 foi inserido.")
	}
}
