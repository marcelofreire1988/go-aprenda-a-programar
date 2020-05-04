//Calcule o tempo de uma viagem de carro. Pergunte a distância a percorrer e a velocidade média esperada para a viagem.

package main

import (
	"fmt"
)

func main() {
	var distancia, velocidadeMedia, tempoTotal float64
	fmt.Println("Digite a distancia em KM:")
	fmt.Scan(&distancia)
	fmt.Println("Digite o valor da velocidade média:")
	fmt.Scan(&velocidadeMedia)
	tempoTotal = distancia / velocidadeMedia
	fmt.Println("O tempo total de viagem será de:", tempoTotal, "horas")

}
