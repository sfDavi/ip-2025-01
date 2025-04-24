package main

import "fmt"

func main() {
	var vetor1 [10]int
	var vetor2 [5]int

	fmt.Println("Digite separdamente 10 números inteiros para o primeiro vetor:")
	for i := 0; i < 10; i++ {
		fmt.Printf("%dº número: ", i+1)
		fmt.Scan(&vetor1[i])
	}

	fmt.Println("\nDigite separdamente 5 números inteiros para o segundo vetor:")
	for i := 0; i < 5; i++ {
		fmt.Printf("%dº número: ", i+1)
		fmt.Scan(&vetor2[i])
	}

	soma := 0
	for _, num2 := range vetor2 {
		soma += num2
	}

	var vetorPares []int
	var vetorImpares []int

	for _, num1 := range vetor1 {
		if num1%2 == 0 {
			vetorPares = append(vetorPares, num1+soma)
		} else {
			vetorImpares = append(vetorImpares, num1+soma)
		}
	}

	fmt.Println("\nPrimeiro vetor resultante:", vetorPares)
	fmt.Println("Segundo vetor resultante:", vetorImpares)
}
