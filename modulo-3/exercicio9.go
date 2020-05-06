//- Crie um programa que utilize a declaração switch, onde o switch statement seja uma variável do tipo string com identificador "esporteFavorito".

package main

import (
	"fmt"
)

func main() {

	esportePredileto := "volei"

	switch esportePredileto {
	case "basquete":
		fmt.Println("Basquete é legal, mas nem tento!")
	case "volei":
		fmt.Println("Assita haikyuu! melhor anime de volei!")
	case "futebol":
		fmt.Println("Sou péssimo")

	default:
		fmt.Println("Gosto mesmo é de dormir")
	}

}
