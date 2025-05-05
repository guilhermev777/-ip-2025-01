package main
import "fmt"

func main() {
var soma float64 
numerador:=100.0
termos:= 20

for i := 0; i < termos ; i++{
    fatorial:=1.0
    for j:=1; j<=i; j++{
        fatorial *= float64(j)
    }
    termo:= numerador / fatorial
  soma+=termo
  
  numerador--
}


                 
 fmt.Printf("A soma dos 20 primeiros termos é: %.10f\n", soma)
}