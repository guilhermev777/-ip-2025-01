package main

import (
	"fmt"
)

func main() {
	var tipoConsumidor string
	var consumo float64

	package main

	import (
		"fmt"
	)
	
	func main() {
		var tipoConsumidor string
		var consumo float64
	
		// Solicitar dados do usuário
		fmt.Println("SANEAGO - Cálculo da Conta de Água")
		fmt.Println("Tipos de consumidor: (R)esidencial, (C)omercial, (I)ndustrial")
		fmt.Print("Digite o tipo de consumidor: ")
		fmt.Scan(&tipoConsumidor)
		fmt.Print("Digite o consumo em m³: ")
		fmt.Scan(&consumo)
	
		// Calcular valor da conta
		var valorConta float64
	
		switch tipoConsumidor {
		case "R", "r":
			valorConta = 5.00 + (0.05 * consumo)
			fmt.Printf("\nCONTA RESIDENCIAL\n")
		case "C", "c":
			if consumo <= 80 {
				valorConta = 500.00
			} else {
				valorConta = 500.00 + (0.25 * (consumo - 80))
			}
			fmt.Printf("\nCONTA COMERCIAL\n")
		case "I", "i":
			if consumo <= 100 {
				valorConta = 800.00
			} else {
				valorConta = 800.00 + (0.04 * (consumo - 100))
			}
			fmt.Printf("\nCONTA INDUSTRIAL\n")
		default:
			fmt.Println("\nTipo de consumidor inválido!")
			return
		}
	
		// Exibir resultados
		fmt.Printf("Consumo: %.2f m³\n", consumo)
		fmt.Printf("Valor da conta: R$ %.2f\n", valorConta)
	}
	fmt.Println("Tipos de consumidor: (R)esidencial, (C)omercial, (I)ndustrial")
	fmt.Print("Digite o tipo de consumidor: ")
	fmt.Scan(&tipoConsumidor)
	fmt.Print("Digite o consumo em m³: ")
	fmt.Scan(&consumo)

	
	var valorConta float64

	switch tipoConsumidor {
	case "R", "r":
		valorConta = 5.00 + (0.05 * consumo)
		fmt.Printf("\nCONTA RESIDENCIAL\n")
	case "C", "c":
		if consumo <= 80 {
			valorConta = 500.00
		} else {
			valorConta = 500.00 + (0.25 * (consumo - 80))
		}
		fmt.Printf("\nCONTA COMERCIAL\n")
	case "I", "i":
		if consumo <= 100 {
			valorConta = 800.00
		} else {
			valorConta = 800.00 + (0.04 * (consumo - 100))
		}
		fmt.Printf("\nCONTA INDUSTRIAL\n")
	default:
		fmt.Println("\nTipo de consumidor inválido!")
		return
	}

	
	fmt.Printf("Consumo: %.2f m³\n", consumo)
	fmt.Printf("Valor da conta: R$ %.2f\n", valorConta)
}