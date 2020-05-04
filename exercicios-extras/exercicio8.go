//Faça agora o contrário, de Fahrenheit para Celsius. (F - 31) / 1.8

package main

import (
	"fmt"
)

func main() {
	var temperaturaC, temperaturaF float64
	fmt.Println("Digite a temperatura em Fahrenheit:")
	fmt.Scan(&temperaturaF)
	temperaturaC = (temperaturaF - 32) / 1.8
	fmt.Println("A temperatura em ° Celcius será de °", temperaturaC)
}
