//Solicite o preço de uma mercadoria e o percentual de desconto. Exiba o valor do desconto e o preço a pagar.

package main

import (
	"fmt"
)

func main() {
	var valorMercadoria, precoFinal, desconto, valorDesconto float64
	fmt.Println("Insira o valor da marcadoria R$: ")
	fmt.Scan(&valorMercadoria)
	fmt.Println("Insiva a porcetagem  do desconto")
	fmt.Scan(&desconto)
	valorDesconto = valorMercadoria * desconto / 100
	precoFinal = valorMercadoria - valorDesconto
	fmt.Println("O preço final da mercadoria com o desconto ficou de R$", precoFinal)
}
