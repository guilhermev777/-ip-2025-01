package main
import "fmt"

func main() {
    
 var salcar float64
 
 fmt.Print("Insira o salário de carlos: ")
 fmt.Scan(&salcar)
 
  saljo := salcar / 3
 fmt.Printf("\nEntão o João ganha %.f reais\n", saljo)
 
 meses := 0
    investimentoJoao := saljo  
    investimentoCarlos := salcar
    
    
    for investimentoCarlos >= investimentoJoao {
        meses++
        investimentoJoao *= 1.05  
        investimentoCarlos *= 1.02  
    }
 fmt.Printf("serão necessários %d meses para os ganhos do joão se igualar aos ganhos do Carlos  ", meses)
}