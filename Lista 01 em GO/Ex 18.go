package main

import "fmt"

func main() {

	var valor_ini, razao, num_ele, soma float64
	var i float64

	fmt.Println("Digite o valor inicial da progressão aritmética")
	fmt.Scan(&valor_ini)
	fmt.Println("Digite a razão da progressão artitmética")
	fmt.Scan(&razao)
	fmt.Println("Digite o número de elementos da progressão")
	fmt.Scan(&num_ele)

	soma = 0

	for i = 1; i <= num_ele; i++ {
		soma = soma + razao*(i-1) + valor_ini
	}

	fmt.Printf("O VALOR DA SOMA É %.0f", soma)
}
