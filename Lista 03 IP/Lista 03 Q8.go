package main
import "fmt"

func main() {
 soma:= 0 
 
 for i:= 1; i <= 20; i++ {
     fmt.Println(i)
     soma += i
 }
 
  
  fmt.Println("\nA soma de todos os números são:",soma)
  
  
}