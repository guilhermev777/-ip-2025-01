package main
import "fmt"

func main() {
var soma float64 
 var n int
 
 fmt.Scan(&n)
 
 for i := 0 ; i <= n ; i++{
     
     
     termo := float64(1000 - 3*i)  /float64(i + 1)
     soma += termo
  
   
 }
 fmt.Print(soma)
}