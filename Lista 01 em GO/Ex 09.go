package main

import "fmt"

func main() {

	var a, b, c float64
	var delta float64

	// ENTRADA

	fmt.Println("Digite o valor do coeficiente A")
	fmt.Scan(&a)
	fmt.Println("Digite o valor do coeficiente B")
	fmt.Scan(&b)
	fmt.Println("Digite o valor do coeficiente C")
	fmt.Scan(&c)

	// PROCESSAMENTO

	delta = (b * b) - (4 * a * c)

	// SAIDA

	fmt.Printf("O valor de delta E = %.2f", delta)

}
