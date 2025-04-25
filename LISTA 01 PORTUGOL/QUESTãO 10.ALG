Algoritmo "Questão 10"


Var
A,B,C,D,det:REAL

Inicio
escreva("Insira o valor da linha 01 e coluna 01: ")
leia(A)

escreva("Insira o valor da linha 01 e coluna 02: ")
leia(B)

escreva("Insira o valor da linha 02 e coluna 01: ")
leia(C)

escreva("Insira o valor da linha 02 e coluna 02: ")
leia(D)

det <- (A*D) - (C*B)

escreva("O valor do determinante é: ", det)

Fimalgoritmo