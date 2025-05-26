package main
import "fmt"

func main() {
    var altura [10] float64
    var soma float64
    media:= soma / 10
    
    for i:=0;i<10;i++{
        fmt.Printf("Insira a altura do atleta %d:",i+1)
        fmt.Scan(&altura[i])
         soma+= altura[i]
       
    }
 
  for i:=0;i<10;i++{
      if altura[i] > media{
          fmt.Print(altura[i])
      }
     
  } 
}