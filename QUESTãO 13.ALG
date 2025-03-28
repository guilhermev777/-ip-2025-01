Algoritmo "questão 13"
 //Em um curso de mestrado as avaliações dos alunos no histórico escolar aparecem em forma de conceito.
//O regulamento do mestrado indica que um professor pode avaliar seus alunos com notas convencionais de
//zero a dez, mas precisa repassar à secretaria do curso a avaliação em termos de conceito. Nesse caso, a
//seguinte tabela de conversão deve ser usada pelo professor:
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
