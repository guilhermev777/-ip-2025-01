package main

import (
	"fmt"
)

func main() {
	var diaSemana, tipoDVD int
	var precoNormal float64

	

	fmt.Println("Dias da semana: 1-Domingo, 2-Segunda, 3-Terça, 4-Quarta, 5-Quinta, 6-Sexta, 7-Sábado")
	fmt.Print("Digite o dia da semana (1-7): ")
	fmt.Scan(&diaSemana)
	
	fmt.Println("Tipos de DVD: 1-Comum, 2-Lançamento")
	fmt.Print("Digite o tipo de DVD (1-2): ")
	fmt.Scan(&tipoDVD)
	
	fmt.Print("Digite o preço normal do aluguel: R$ ")
	fmt.Scan(&precoNormal)

	
	precoFinal := precoNormal

	/
	switch diaSemana {
	case 2, 3, 5: // Segunda, terça, quinta
		precoFinal *= 0.6 // 40% de desconto
	}

	
	if tipoDVD == 2 { // Lançamento
		precoFinal *= 1.15 // 15% de acréscimo
	}

	
	fmt.Printf("\nResumo do Aluguel:\n")
	fmt.Printf("Dia da semana: %d\n", diaSemana)
	fmt.Printf("Tipo de DVD: %s\n", map[int]string{1: "Comum", 2: "Lançamento"}[tipoDVD])
	fmt.Printf("Preço normal: R$ %.2f\n", precoNormal)
	fmt.Printf("Preço final: R$ %.2f\n", precoFinal)
}