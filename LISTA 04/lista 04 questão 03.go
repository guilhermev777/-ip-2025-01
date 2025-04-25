package main
import "fmt"

func main() {
 var numeros [10] int
 var pares [] int 
 var impares [] int
 
 somapares := 0 
 
 fmt.Println("Insira 10 números inteiros: ")
 
 for i := 0 ; i < 10 ; i++ {
     fmt.Printf("Número %d: ", i+1)
     fmt.Scan(&numeros[i])
     
 }
 for _, num := range numeros {
     if num % 2 == 0 { 
         pares = append(pares,num)
         somapares += num
     
 } else { 
     impares = append(impares,num)
        }
    }
fmt.Println("Números pares: ", pares)
fmt.Println("soma dos números pares: ", somapares)
fmt.Println("Números impares: ", impares)
fmt.Println("quantidade de números impares: ", len(impares))
    }
}