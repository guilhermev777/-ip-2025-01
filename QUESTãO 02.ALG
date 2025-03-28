Algoritmo "Questão 02"

var
//ingressos
vip,vig,via,vic:real
pubtot:real
rendatot:real


Inicio

escreva("qual foi o publico total: ")
leia(pubtot)

escreva("qual foi a percentagem de pessoas que compraram ingresso na categoria Popular: ")
leia(vip)

escreva("qual foi a percentagem de pessoas que compraram ingresso na categoria Geral: ")
leia(vig)

escreva("qual foi a percentagem de pessoas que compraram ingresso na categoria Arquibancada :")
leia(via)

escreva("qual foi a percentagem de pessoas que compraram ingresso na categoria Cadeiras. :")
leia(vic)

rendatot <- (vip*pubtot*1) + (vig*5*pubtot) + (via*10*pubtot) + (vic*20*pubtot)

escreva("a renda total foi de", rendatot ," reais")
Fimalgoritmo