Algoritmo "Questão 01"
//Fazer um algoritmo que calcule a média aritmética das três notas de um aluno e mostre, além do valor
//da média, uma mensagem de "APROVADO", caso a média seja igual ou superior a seis, ou a mensagem
//"REPROVADO", caso contrário.

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

escreval("sua média é: ",media)

   se (media >= 6) entao
      escreva("Você foi aprovado!")
   senao
      escreva("Você foi reprovado!")

      fimse
      
Fimalgoritmo