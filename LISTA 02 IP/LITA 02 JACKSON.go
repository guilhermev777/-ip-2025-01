package main
import "fmt"

func main() {
    var n int
    fmt.Print("digite um numero: ")
  fmt.Scan(&n)
  
  if (n > 0 ) {
      fmt.Print("O número é poitivo")
 } else {
    if (n < 0 )  {
    fmt.Print("O número é negativo")
    } else {
        fmt.Print("O número é nulo")
    }   
    }
  }