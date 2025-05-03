package main

import "fmt"

func main() {
	var n = []int{1, 2, 3, 4, 5}
	fmt.Println(soma(n))
}

func soma(n []int) int {
	if len(n) == 0 {
		return 0
	}
	return n[0] + soma(n[1:])
}
