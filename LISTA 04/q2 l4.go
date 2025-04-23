package main

import "fmt"

func main() {
	// Declaração dos vetores
	var vetor1 [10]int
	var vetor2 [5]int
	var paresSoma []int
	var imparesSoma []int

	// Preenchimento do primeiro vetor
	fmt.Println("Digite os 10 números do primeiro vetor:")
	for i := 0; i < 10; i++ {
		fmt.Printf("Elemento %d: ", i+1)
		fmt.Scan(&vetor1[i])
	}

	// Preenchimento do segundo vetor
	fmt.Println("\nDigite os 5 números do segundo vetor:")
	for i := 0; i < 5; i++ {
		fmt.Printf("Elemento %d: ", i+1)
		fmt.Scan(&vetor2[i])
	}

	// Calcula a soma total do segundo vetor
	somaSegundo := 0
	for _, num := range vetor2 {
		somaSegundo += num
	}

	// Processa os números do primeiro vetor
	for _, num := range vetor1 {
		if num%2 == 0 { // Se for par
			paresSoma = append(paresSoma, num+somaSegundo)
		} else { // Se for ímpar
			imparesSoma = append(imparesSoma, num+somaSegundo)
		}
	}

	// Exibe os resultados
	fmt.Println("\nPrimeiro vetor resultante (pares + soma do segundo vetor):", paresSoma)
	fmt.Println("Segundo vetor resultante (ímpares + soma do segundo vetor):", imparesSoma)
}
