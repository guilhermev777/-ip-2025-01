ppackage main
import "fmt"

func main() {
 var a int
 fmt.Print("Digite o número do ângulo em radianos: ")
 fmt.Scan(&a)
 
senoA:= a - (a*a*a / 6) + (a*a*a*a*a / 120) - (a*a*a*a*a*a*a / 5040)
 fmt.Print(senoA)
}