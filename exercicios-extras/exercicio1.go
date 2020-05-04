//Faça um programa que peça dois números inteiros e imprima a soma desses dois números

package main

import (
	"fmt"
)

var num1 int
var num2 int

func main() {
	fmt.Println("------ Bem vindo ------")
	fmt.Println("Digite o primeiro número: ")
	fmt.Scanln(&num1)
	fmt.Println("Digite o segundo número")
	fmt.Scanln(&num2)
	resultado := num1 + num2
	fmt.Println("O resultado é:", resultado)

}
