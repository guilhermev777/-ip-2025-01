package main
import "fmt"

func main() {
 var b, n int
  
fmt.Print("Escreva a base: ")
fmt.Scan(&b)

fmt.Print("Escreva o expoente: ")
fmt.Scan(&n)
 
 resultado := 1
 for i := 0; i < n; i++{
     resultado *= b
 }
 fmt.Print(resultado)
}