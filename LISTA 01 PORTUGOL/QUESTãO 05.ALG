algoritmo "quest�o 05"

var
tipocliente:caractere
contacliente: inteiro
consumoagua,custot:real
inicio

escreva("Escreva o n�mero da sua conta: ")
leia(contacliente)

escreva("digite seu consumo de �gua: ")
leia(consumoagua)

escreva("Qual o tipo de cliente(R,I,C): ")
leia(tipocliente)

   se (tipocliente=R) ent�o
      custot <- 0.5 * consumoagua + 5
      
   senao
       se (tipocliente=I) entao
    
      custot <- 0.25 * consumoagua + 480

       senao
            se(tipocliente=C) entao
      custot <- 0.04 * consumoagua + 796
      
             fimse
           fimse
          fimse
   escreva("O valor da conta �: ",custot)
    
    
fimalgoritmo
