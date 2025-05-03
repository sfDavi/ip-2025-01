package main

import "fmt"

func main() {
	var n int
	fmt.Print("Digite um número: ")
	fmt.Scan(&n)
	fmt.Printf("Fatorial de %d é %d\n", n, fatorial(n))
}

func fatorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * fatorial(n-1)
}
