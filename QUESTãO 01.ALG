Algoritmo "Quest�o 01"
//Fazer um algoritmo que calcule a m�dia aritm�tica das tr�s notas de um aluno e mostre, al�m do valor
//da m�dia, uma mensagem de "APROVADO", caso a m�dia seja igual ou superior a seis, ou a mensagem
//"REPROVADO", caso contr�rio.

var
n1,n2,n3,media:real


Inicio
escreva("digite a nota do bimestre 1: ")
leia(n1)

escreva("digite a nota do bimestre 2: ")
leia(n2)

escreva("digite a nota do bimestre 3: ")
leia(n3)

media <- (n1+n2+n3) / 3

escreval("sua m�dia �: ",media)

   se (media >= 6) entao
      escreva("Voc� foi aprovado!")
   senao
      escreva("Voc� foi reprovado!")

      fimse
      
Fimalgoritmo