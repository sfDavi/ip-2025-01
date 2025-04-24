package main

import (
    "fmt"
)

func main() {
    var numero int

    fmt.Print("Digite um número inteiro: ")
    fmt.Scanln(&numero)

    if numero%5 == 0 {
        fmt.Println("O número inserido é divisível por 5.")
    } else {
        fmt.Println("O número inserido NÃO é divisível por 5.")
    }
}