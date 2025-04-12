package main
import "fmt"

func main() {
    var n int
  fmt.Print("Insira um número inteiro positivo de três digitos: ")
  fmt.Scan(&n)
  
  if n < 100 && n > 999 { 
  fmt.Print("Número inválido ")     
  } else { 
      dezena := (n / 10) % 10
     fmt.Print("O algarismo das dezenas é: ",dezena)   
  }
}