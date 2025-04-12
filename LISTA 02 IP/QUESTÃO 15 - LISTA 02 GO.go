package main

import (
	"fmt"
)

func main() {
	var dia, mes, ano int

	
	fmt.Print("Digite o dia: ")
	fmt.Scan(&dia)
	
	fmt.Print("Digite o mês: ")
	fmt.Scan(&mes)
	
	fmt.Print("Digite o ano: ")
	fmt.Scan(&ano)

	
	nomesMeses := []string{
		"janeiro", "fevereiro", "março", "abril",
		"maio", "junho", "julho", "agosto",
		"setembro", "outubro", "novembro", "dezembro",
	}

	
	fmt.Printf("%d de %s de %d\n", dia, nomesMeses[mes-1], ano)
}