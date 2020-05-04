//Faça um programa que calcule o aumento de um salário. Ele deve solicitar o valor do salário e a porcentagem do aumento. Exiba o valor do aumento e do novo salário.
package main

import (
	"fmt"
)

func main() {
	var salarioAtual, porcetagemAumento float64
	fmt.Println("Digite o seu salário: ")
	fmt.Scan(&salarioAtual)
	fmt.Println("Digite a porcetagem de aumento desejada: ")
	fmt.Scan(&porcetagemAumento)

	novoSalario := salarioAtual * porcetagemAumento / 100
	fmt.Println("Seu novo salário é de: ", novoSalario)

}
