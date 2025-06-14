package main

import "fmt"

func main() {
	
	var numeros [10]int
	var contador int

	
	fmt.Println("Digite 10 números inteiros:")
	for i := 0; i < 10; i++ {
		fmt.Printf("Número %d: ", i+1)
		fmt.Scan(&numeros[i])
	}

	
	fmt.Println("\nNúmeros maiores que 50 e suas posições:")
	for indice, valor := range numeros {
		if valor > 50 {
			fmt.Printf("Número: %d, Posição: %d\n", valor, indice)
			contador++
		}
	}

	
	if contador == 0 {
		fmt.Println("Não existem números maiores que 50 no vetor.")
	}
}
