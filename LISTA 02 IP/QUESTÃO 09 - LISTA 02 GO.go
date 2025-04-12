package main
import "fmt"

func main() {
    var vc, vv float64
  fmt.Print("Insira o valor de compra: ")
  fmt.Scan(&vc)
  
  if vc < 10.00 {
      vv = vc * 1.7
    
  } else if vc >= 10.00 && vc < 30.00 {
      vv = vc * 1.5
    
} else if vc >= 30.00 && vc < 50.00 {
      vv = vc * 1.4
      
} else  if vc >= 50.00 {
      vv = vc * 1.3
 }
  fmt.Print("O valor de venda será: ",vv )
}