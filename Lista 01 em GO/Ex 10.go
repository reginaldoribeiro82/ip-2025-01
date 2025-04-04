package main

import "fmt"

func main() {

	var a, b, c, d float64
	var delta float64

	// ENTRADA

	fmt.Println("Digite o valor do valor a")
	fmt.Scan(&a)
	fmt.Println("Digite o valor do valor b")
	fmt.Scan(&b)
	fmt.Println("Digite o valor do valor c")
	fmt.Scan(&c)
	fmt.Println("Digite o valor do valor d")
	fmt.Scan(&d)

	// PROCESSAMENTO

	delta = (a * d) - (b * c)

	// SAIDA

	fmt.Printf("O VALOR DO DETERMINANTE É = %.2f", delta)

}
