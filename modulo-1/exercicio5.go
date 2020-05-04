package main

import (
	"fmt"
)

//exercicio passado
var x exercicio4

type exercicio4 int

//novo exercicio
var y int

func main() {
	//exercicio passado
	fmt.Printf("%v\n, %T\n", x, x)
	x = 42
	fmt.Println(x)

	// novo exercicio
	y = int(x)
	fmt.Printf("%v\n, %T\n", y, y)

}
