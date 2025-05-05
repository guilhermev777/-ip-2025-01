package main
import "fmt"

func main() {
 rows, cols := 10,10
    
 for i:=0;i<rows;i++{
    for j:=0;j<cols;j++{
        
       if j <= i {
        fmt.Printf("[%d][%d] ",i, j)
       }
        
 }   
   fmt.Println()
 }
 
}