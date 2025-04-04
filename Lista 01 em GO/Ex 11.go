package main

import "fmt"

func main() {

	var num uint64
	var divisao_3, divisao_5 uint64

	fmt.Println("Digite um numero inteiro")
	fmt.Scan(&num)

	divisao_3 = num % 3
	divisao_5 = num % 5

	if divisao_3 == 0 && divisao_5 == 0 {
		fmt.Println("O numero é divisivel")
	} else {
		fmt.Println("O numero não é divisivel")
	}
}
