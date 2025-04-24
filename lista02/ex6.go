package main

import (
    "fmt"
)

func main() {
    var a, b int

    fmt.Print("Digite o número inteiro A: ")
    fmt.Scanln(&a)

    fmt.Print("Digite o número inteiro B: ")
    fmt.Scanln(&b)

    if b == 0 {
        fmt.Println("Divisão por zero não é permitida.")
    } else if a%b == 0 {
        fmt.Printf("O número %d é divisível por %d.\n", a, b)
    } else {
        fmt.Printf("O número %d NÃO é divisível por %d.\n", a, b)
    }
}