Algoritmo "Questão 07"

Var
r,h,atot,vc,p:real

Inicio
escreva("Digite a altura da lata em metro: ")
leia(h)

escreva("Digite o raio da lata em metro: ")
leia(r)

p <-3.14159: const

Atot <- 2*(p*r^2) + 2*p*r*h

vc <- atot*100

escreval("O valor de custo é R$: ", vc)


Fimalgoritmo