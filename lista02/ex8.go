package main

import (
    "fmt"
)

func main() {
    var numero int

    fmt.Print("Digite um número inteiro: ")
    fmt.Scanln(&numero)

    if numero > 20 && numero < 90 {
        fmt.Println("O número está na faixa entre 20 e 90 (exclusivo).")
    } else {
        fmt.Println("O número NÃO está na faixa entre 20 e 90 (exclusivo).")
    }
}