Algoritmo "quest�o 13"
 //Em um curso de mestrado as avalia��es dos alunos no hist�rico escolar aparecem em forma de conceito.
//O regulamento do mestrado indica que um professor pode avaliar seus alunos com notas convencionais de
//zero a dez, mas precisa repassar � secretaria do curso a avalia��o em termos de conceito. Nesse caso, a
//seguinte tabela de convers�o deve ser usada pelo professor:
var
nota: real

Inicio
 escreva("Digite sua nota: ")
 leia(nota)
 
    se (nota>=9) entao
      escreva("A")
    senao
    
       se (nota<9) e (nota>=7.5) entao
         escreva("B")
       senao
       
          se (nota>6) e (nota < 7.5) entao
          escreva("C")
          senao
          escreva("D")
          
          fimse
         fimse
        fimse
Fimalgoritmo
