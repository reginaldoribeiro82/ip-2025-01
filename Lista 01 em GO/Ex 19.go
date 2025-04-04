package main

import "fmt"

func main() {

	var num, soma, i float64

	fmt.Println("Digite um numero inteiro positivo")
	fmt.Scan(&num)
	soma = 0
	if num >= 1 {
		for i = 1; i <= num; i++ {
			soma = soma + (1 / i)
		}
		fmt.Printf("O valor de S é %.6f", soma)
	} else {
		fmt.Println("Número inválido")
	}
}
