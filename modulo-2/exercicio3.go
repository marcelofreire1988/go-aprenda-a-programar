package main

import (
	"fmt"
)

func main() {
	x := 88

	fmt.Printf("O número em decimal é: %d\n", x)
	fmt.Printf("O número em binário é: %b\n", x)
	fmt.Printf("O número em hexadecimal é: %#x\n", x)

	y := x << 1
	fmt.Println("Deslocando o valor da variável X uma casa para a esquerda\n")
	fmt.Printf("O número em decimal é: %d\n", y)
	fmt.Printf("O número em binário é: %b\n", y)
	fmt.Printf("O número em hexadecimal é: %#x", y)

}
