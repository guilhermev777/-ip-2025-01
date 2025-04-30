package main
import "fmt"

func main() {

h:= 0.0
 for i:=1; i<=50 ;i++ {
  numerador:= 2.0*float64(i) - 1.0
  denominador:= float64(i)
  
  termo:= numerador / denominador
  
  h += termo
  
     
 } 
 
 fmt.Print(h)
 
}
