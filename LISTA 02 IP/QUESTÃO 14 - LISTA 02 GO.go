package main

import (
	"fmt"
)

func main() {
	var precoFabrica float64
	var temArCondicionado, temPinturaMetalica, temVidroEletrico, temDirecaoHidraulica string

	
	fmt.Print("Digite o preço de fábrica do carro: R$ ")
	fmt.Scan(&precoFabrica)

	
	fmt.Print("Deseja ar condicionado? (s/n): ")
	fmt.Scan(&temArCondicionado)

	fmt.Print("Deseja pintura metálica? (s/n): ")
	fmt.Scan(&temPinturaMetalica)

	fmt.Print("Deseja vidro elétrico? (s/n): ")
	fmt.Scan(&temVidroEletrico)

	fmt.Print("Deseja direção hidráulica? (s/n): ")
	fmt.Scan(&temDirecaoHidraulica)

	
	precoFinal := precoFabrica

	if temArCondicionado == "s" || temArCondicionado == "S" {
		precoFinal += 1750.00
	}

	if temPinturaMetalica == "s" || temPinturaMetalica == "S" {
		precoFinal += 800.00
	}

	if temVidroEletrico == "s" || temVidroEletrico == "S" {
		precoFinal += 1200.00
	}

	if temDirecaoHidraulica == "s" || temDirecaoHidraulica == "S" {
		precoFinal += 2000.00
	}


	fmt.Printf("\nPreço final do carro: R$ %.2f\n", precoFinal)
}