Algoritmo "questão 14"

var
Aresta,h,v,ab: real

Inicio
 escreva("Digite o tamanho da aresta em metros: ")
 leia(aresta)
 
 escreva("Digite o tamanho da altura da pirâmide em metros: ")
 leia(h)
 
 ab <- 3 * aresta^2 * 3^0.5 / 2
 
 v <- (1/3) * ab * h
 
 escreval("O volume da pirâmide é: ", v, "metros cúbicos")
 
 
 
 
Fimalgoritmo
