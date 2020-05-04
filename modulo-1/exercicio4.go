package main

import (
	"fmt"
)

var x exercicio4

type exercicio4 int

func main() {
	fmt.Printf("%v\n, %T\n", x, x)
	x = 42
	fmt.Println(x)
}
