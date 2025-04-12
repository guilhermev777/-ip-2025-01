package main
import "fmt"

func main() {
    var x float64
  fmt.Print("Insira o valor de x: ")
  fmt.Scan(&x)
  
  if x == 2  {
 fmt.Print("Não é possível realizar divisão por zero ")
  } else {
  
 y := 8 / (2-x)
 fmt.Print("Valor de f(x): ",y)
}
}