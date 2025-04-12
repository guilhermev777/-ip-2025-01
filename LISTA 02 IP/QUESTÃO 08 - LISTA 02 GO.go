package main
import "fmt"

func main() {
    n int
    
  fmt.Println("Digite um número: ")
  fmt.Scan(&n)
  
  if n >= 20 && n <= 90 {
      fmt.Print("O número está entre 20 e 90")
  } else { 
      fmt.Print("O número não está entre 20 e 90")
}
}