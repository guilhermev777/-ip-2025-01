package main

import "fmt"

func main() {
    // Declaração das variáveis
    var a, b, c int
    
    // Leitura dos valores
    fmt.Print("Digite o primeiro valor: ")
    fmt.Scan(&a)
    
    fmt.Print("Digite o segundo valor: ")
    fmt.Scan(&b)
    
    fmt.Print("Digite o terceiro valor: ")
    fmt.Scan(&c)
    
    // Inicialização das variáveis para ordenação
    menor := a
    inter := b
    maior := c
    
    // Encontrando o menor valor
    if b < menor {
        menor = b
    }
    if c < menor {
        menor = c
    }
    
    // Encontrando o maior valor
    if b > maior {
        maior = b
    }
    if a > maior {
        maior = a
    }
    
    // Encontrando o valor intermediário
    // O intermediário é a soma dos três menos o maior e o menor
    inter = a + b + c - menor - maior
    
    // Imprimindo os valores em ordem crescente
    fmt.Printf("Valores em ordem crescente: %d, %d, %d\n", menor, inter, maior)
}