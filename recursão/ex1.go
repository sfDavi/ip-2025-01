package main

import "fmt"

func main() {
	var x, n int
	fmt.Scanln(&x, &n)
	fmt.Printf("%d elevado a %d é %d\n", x, n, pot(x, n))
}

func pot(x int, n int) int {
	if n == 0 {
		return 1
	}
	return x * pot(x, n-1)
}
