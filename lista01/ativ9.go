package main

import (
	"fmt"
)

func main() {
	var a, b, c float64

	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)

	delta := (b * b) - (4 * a * c)

	fmt.Printf("O VALOR DE DELTA E = %.2f\n", delta)
}
