algoritmo "Custo_Energia"
var
    salarioMinimo, qtdKw, custoKw, custoConsumo, custoDesconto: real
inicio
    escreva("Digite o salário mínimo: ")
    leia(salarioMinimo)
    
    escreva("digite a quantidade de energia gasta: ")
    leia(qtdkw)

    custoKw <- (0.7 * salarioMinimo) / 100
    custoConsumo <- custoKw * qtdKw
    custoDesconto <- custoConsumo * 0.9

    escreval("Custo por kW: R$ ", custoKw)
    
    escreval("Custo do consumo: R$ ", custoConsumo)
    
    escreval("Custo com desconto: R$ ", custoDesconto)
    
    
fimalgoritmo
