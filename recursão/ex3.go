package main

import "fmt"

func main() {
	var n = []int{1, 2, 3, 4, 5}
	fmt.Println(inverte(n, len(n)))
}

func inverte(n []int, l int) []int {
	if l == 0 {
		return []int{}
	}
	return (append([]int{n[l-1]}, inverte(n, l-1)...))
}
