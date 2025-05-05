package main
import "fmt"

func main() {

var soma float64 = 0
 
 for i := 1; i <= 37 ; i++{
   n1:= (37 - i + 1)
   n2:= (37 - i + 2)
   
   termo :=float64(n1 * n2) / float64(i)
   
   soma += termo 
 }
 fmt.Print(soma)
}