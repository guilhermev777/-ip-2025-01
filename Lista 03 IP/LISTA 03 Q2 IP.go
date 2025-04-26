package main
import "fmt"

func main() {

  soma := 0
  quantpares := 0
  
  fmt.Println("Números de 50 a 70\n")
  
  for n := 50 ; n <= 70; n++ {
      
      if n % 2 == 0 {
          fmt.Printf("%d ",n)
          soma = soma + n
          quantpares++
      }
  }
  
  media := 0.0
  if quantpares > 0 {
      media = float64(soma) / float64(quantpares)
  }
  
 fmt.Printf("\nSoma dos pares: %d\n", soma)
    fmt.Printf("Média aritmética: %.f\n", media)  
    fmt.Printf("Quantidade de pares: %d\n", quantpares)
  
  
}