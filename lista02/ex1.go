package main

import (
    "fmt"
)

func main() {
    var numero int

    fmt.Print("Digite um número inteiro: ")
    fmt.Scanln(&numero)

    if numero%2 == 0 {
        fmt.Println("O número inserido é PAR.")
    } else {
        fmt.Println("O número inserido é ÍMPAR.")
    }
}