package main
import "fmt"

func main() {
 var n int 
 
  fmt.Print("Digite o número de alunos: ")
  fmt.Scan(&n)
  
  var reprovados,aprovados,exame  int
  var mediaclasse float64
 
 for i:= 1; i <= n; i++ {
     var nota1, nota2 float64
     fmt.Printf("\naluno %d: ",i)
     
     fmt.Printf("\nnota 1: ")
     fmt.Scan(&nota1)
     
     fmt.Printf("nota 2: ")
     fmt.Scan(&nota2)
    
    media := (nota1 + nota2)/2
    mediaclasse += media
    
     fmt.Printf("A média do aluno %d é igual a %.f", i, media) 
 
 
 if media <= 3 {
     fmt.Println("\nReprovado")
     reprovados++
     } else if media > 3 && media <= 7 {
         fmt.Println("\nExame") 
         exame++
     } else {
         fmt.Println("\nAprovado")
         aprovados++
     }
 
 }
 mediaclasse =  mediaclasse / float64(n)
 
 fmt.Printf("Total de aprovados: %d\n", aprovados)
	fmt.Printf("Total em exame: %d\n", exame)
	fmt.Printf("Total de reprovados: %d\n", reprovados)
	fmt.Printf("Média da classe: %.1f\n", mediaclasse)
  
  
}