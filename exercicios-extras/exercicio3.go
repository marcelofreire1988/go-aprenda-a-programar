//Escreva um programa que leia a quantidade de dias, horas, minutos e segundos do usuário. Calcule o total em segundos.

package main

import (
	"fmt"
)

func main() {
	var dias, horas, minutos int
	var segundos float64
	fmt.Println("Insira um dia: ")
	fmt.Scan(&dias)
	fmt.Println("O dia inserido foi:", dias)
	fmt.Println("Insira uma hora: ")
	fmt.Scan(&horas)
	fmt.Println("A hora inserida foi: ", horas)
	fmt.Println("Insira o minuto: ")
	fmt.Scan(&minutos)
	fmt.Println("O minuto inserido foi:", minutos)
	fmt.Println("Insira os segundos: ")
	fmt.Scan(&segundos)
	fmt.Println("O segundo inserido foi:", segundos)
	total := float64(dias)*24*60*60 + float64(horas)*60*60 + float64(minutos)*60 + segundos
	fmt.Println("O total convertido em segundos é", total)

}
