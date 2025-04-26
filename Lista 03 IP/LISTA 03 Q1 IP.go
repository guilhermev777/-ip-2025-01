package main
import "fmt"
     
func main() {
    var base float64
    var expoente int
 fmt.Print("Escreva a base: ")
 fmt.Scan(&base)
 fmt.Print("Escreva o expoente: ")
 fmt.Scan(&expoente)
 
 if base <= 0 {
     fmt.Print("O número deve ser positivo")
     return
 }
 
 valor := 1.0
 for i := 0; i < expoente; i++ {
     valor = valor * base
 }  
 
 fmt.Println(valor)
  
 
 
}