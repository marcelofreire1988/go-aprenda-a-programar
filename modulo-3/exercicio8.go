//- Crie um programa que utilize a declaração switch, sem switch statement especificado.

package main

import (
	"fmt"
)

func main() {

	x := 1
	switch {
	case x != 1:
		fmt.Println("é isso ai")
	case x > 10:
		fmt.Println("maior")
	case x == 1:
		fmt.Println("igual")
	default:
		fmt.Println("default")
	}
}
