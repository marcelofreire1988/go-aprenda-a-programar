// - Crie um loop utilizando a sintaxe: for {}
// - Utilize-o para demonstrar os anos desde que você nasceu.

package main

import (
	"fmt"
)

func main() {
	anoNascimento := 1988
	anoAtual := 2020
	for {
		if anoNascimento > anoAtual {
			break
		}
		fmt.Println(anoNascimento)
		anoNascimento++
	}
}
