//- Utilizando a solução anterior, adicione as opções else if e else.

package main

import (
	"fmt"
)

func main() {
	x := 12
	if x == 10 {
		fmt.Println("teste de if")

	} else if x < 10 {
		fmt.Println("teste de else if")
	} else {
		fmt.Println("teste de else")
	}

}
