package main

import "fmt"

func main() {
	var horas, minutos, segundos float64
	var min_convert, horas_convert, soma float64

	fmt.Println("Digite o valor em horas")
	fmt.Scan(&horas)
	fmt.Println("Digite o valor em minutos")
	fmt.Scan(&minutos)
	fmt.Println("Digite o valor em segundos")
	fmt.Scan(&segundos)

	min_convert = minutos * 60
	horas_convert = horas * 3600
	soma = segundos + min_convert + horas_convert

	fmt.Printf("O TEMPO EM SEGUNDOS É %.0f", soma)

}
