//Escreva um programa que leia um valor em metros e o exiba convertido em milímetros

package main

import (
	"fmt"
)

var metros float64

func main() {

	fmt.Println("Digite um valor em metros:")
	fmt.Scanln(&metros)
	centimetros := (metros * 100)
	fmt.Println("O valor em centimetros é de:", centimetros)

}
