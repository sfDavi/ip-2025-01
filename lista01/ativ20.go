package main

import (
	"fmt"
)

func main() {
	var h, min, seg int

	fmt.Scanln(&h)
	fmt.Scanln(&min)
	fmt.Scanln(&seg)

	totalSegundos := (h * 3600) + (min * 60) + seg

	fmt.Printf("O TEMPO EM SEGUNDOS E = %d\n", totalSegundos)
}
