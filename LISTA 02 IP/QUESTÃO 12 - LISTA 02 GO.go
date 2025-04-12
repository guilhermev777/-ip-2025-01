package main
import "fmt"

func main() {
    var n int
  fmt.Print("Insira sua idade: ")
  fmt.Scan(&n)
  
  if n <= 2  {
 fmt.Print("RECÉM NASCIDO ")
 
  } else if n > 2 && n <= 11  {
   fmt.Print("CRIANÇA ")
   
  } else if n > 11 && n <= 19  {
   fmt.Print("Adolescente ")
   
  } else if n > 19 && n <= 55  {
   fmt.Print("ADULTO")
   
  } else if n > 55  {
   fmt.Print("IDOSO")
  
}
}