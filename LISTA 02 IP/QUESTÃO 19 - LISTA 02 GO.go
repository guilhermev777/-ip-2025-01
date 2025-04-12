package main

import (
	"fmt"
	"math"
)

func main() {
	var figura int
	var raio, altura float64


	fmt.Println("Escolha a figura:")
	fmt.Println("1 - Cone")
	fmt.Println("2 - Cilindro")
	fmt.Println("3 - Esfera")
	fmt.Scan(&figura)

	
	fmt.Print("Digite o raio: ")
	fmt.Scan(&raio)

	
	if figura == 1 || figura == 2 {
		fmt.Print("Digite a altura: ")
		fmt.Scan(&altura)
	}

	
	switch figura {
	case 1:
		volume := (math.Pi * raio * raio * altura) / 3
		area := math.Pi * raio * (raio + math.Sqrt(raio*raio + altura*altura))
		fmt.Printf("Volume do cone: %.2f\n", volume)
		fmt.Printf("Área do cone: %.2f\n", area)

	case 2: 
		volume := math.Pi * raio * raio * altura
		area := 2 * math.Pi * raio * (raio + altura)
		fmt.Printf("Volume do cilindro: %.2f\n", volume)
		fmt.Printf("Área do cilindro: %.2f\n", area)
		
    case 3: 
		volume := (4 * math.Pi * raio * raio * raio) / 3
		area := 4 * math.Pi * raio * raio
		fmt.Printf("Volume da esfera: %.2f\n", volume)
		fmt.Printf("Área da esfera: %.2f\n", area)

	default:
		fmt.Println("Opção inválida!")
	}
}