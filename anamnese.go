package main

import (
    "fmt"
)

func main() {
    var peso, altura float64

    // Entrada de dados
    fmt.Print("Digite seu peso (kg): ")
    fmt.Scan(&peso)

    fmt.Print("Digite sua altura (m): ")
    fmt.Scan(&altura)

    // Cálculo do IMC
    imc := peso / (altura * altura)

    // Saída do resultado
    fmt.Printf("Seu IMC é: %.2f\n", imc)

    // Classificação
    switch {
    case imc < 18.5:
        fmt.Println("Classificação: Abaixo do peso")
    case imc >= 18.5 && imc < 25:
        fmt.Println("Classificação: Peso normal")
    case imc >= 25 && imc < 30:
        fmt.Println("Classificação: Sobrepeso")
    default:
        fmt.Println("Classificação: Obesidade")
    }
}
