package main
import "fmt"

func main() {
    var reg, sit, vp  int
  fmt.Print("Para qual região vc vai:norte(1), nordeste(2), sul(4) ou centro-oeste(3): ")
  fmt.Scan(&reg)
  
  fmt.Print("Opção 1(Ida e volta) ou Opção 2(somente ida):")
  fmt.Scan(&sit)
  
  if reg == 1 && sit == 1 { 
      vp=900
  } else if reg ==1 && sit==2 {
      vp=500
  } else if reg==2 && sit==1 {
      vp=650
  } else if reg==2 && sit==2 {
      vp=350
  } else if reg==3 && sit==1 {  
      vp=600
  } else if reg==3 && sit==2 {
      vp=350
  } else if reg==4 && sit==1 {
      vp=550
  } else if reg==4 && sit== 2 {
      vp=300
  }
 fmt.Print("PREÇO DA PASSAGEM: ",vp)
}