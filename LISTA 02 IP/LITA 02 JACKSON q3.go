package main
import (
        "fmt"
        "math"
    )

func main() {
    var n float64
    
    if (n >= 0) {
    raizquadrada := math.sqrt(n)
    fmt.Print("A raiz quadrada é ", n, raizquadrada)
    
    } else {
    quadrado := n * n
    fmt.Print("O quadrado é ", n, quadrado)
    }
    
   
    
    
  
}