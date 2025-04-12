package main
import "fmt"

func main() {
    var n1, n2 int
    
    fmt.Print("digite um numero: ")
  fmt.Scan(&n1)
    fmt.Print("digite um numero: ")
  fmt.Scan(&n2)
  
  x := n1 + n2
  if (x > 20 ) {
      z:= x + 8
      fmt.Print(z)
 } else {
    z := x - 5
    fmt.Print(z)
    
    }
  }