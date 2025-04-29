package main
import "fmt"

func main() {
  fmt.Println("Insira números inteiros positivos (0 para sair):  ")
  
  for {
      var num int 
      fmt.Print("Números : ")
      fmt.Scan(&num)
      
      if num <= 0 {
          fmt.Println("PROGRAMA ENCERRADO")
          break
      }
      oquadradoeperfeito:= false
      for i := 0; i*i <= num ; i++ {
          if num == i*i {
              oquadradoeperfeito= true
              break
          }
      }
       if  oquadradoeperfeito {
            fmt.Printf("%d é um quadrado perfeito!\n", num)
        } else {
            fmt.Printf("%d NÃO é um quadrado perfeito.\n", num)
        }
  }
  
}