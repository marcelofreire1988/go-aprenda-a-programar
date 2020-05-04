//Converta uma temperatura digitada em Celsius para Fahrenheit. F = 9*C/5 + 32
package main

import (
	"fmt"
)

func main() {
	var temperaturaCelsius, temperaturaFahrenheit float64
	fmt.Println("Digite a temperatura em ° Celsius: ")
	fmt.Scan(&temperaturaCelsius)
	temperaturaFahrenheit = temperaturaCelsius*9/5 + 32
	fmt.Println("A temperatura em Fahrenheit é de ° :", temperaturaFahrenheit, "Fahrenheit")

}
