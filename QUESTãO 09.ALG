Algoritmo "Questão 09"

Var
A,B,C,delta,x1,x2:real

Inicio

escreva("Digite o valor de A: ")
leia(A)

escreva("Digite o valor de B: ")
leia(B)

escreva("Digite o valor de C: ")
leia(C)

delta <- B^2 - 4 * A * C

x1 <- (-B + delta^0.5) / (2*A)

x2 <- (-B - delta^0.5) / (2*A)

escreval("o resultado de X1 é: ",x1)
escreval("o resultado de X2 é: ",x2)
Fimalgoritmo