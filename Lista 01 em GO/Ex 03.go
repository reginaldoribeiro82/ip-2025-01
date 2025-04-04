package main

import "fmt"

func main() {
	var n1, n2, n3 float64
	var n1_cent, n2_dez float64
	var quadrado float64
	var n_total float64

	fmt.Println("Digite um número (centena)")
	fmt.Scan(&n1)

	fmt.Println("Digite um número (dezena)")
	fmt.Scan(&n2)

	fmt.Println("Digite um número (unidade)")
	fmt.Scan(&n3)
	
	if n1 < 9 && n2 < 9 && n3 < 10  {
		n1_cent = n1 * 100
		n2_dez = n2 * 10
		n_total = n1_cent + n2_dez + n3
		quadrado = n_total * n_total
		fmt.Printf("%.0f%.0f%.0f, %.0f", n1, n2, n3, quadrado)
	}else {
		fmt.Println("DIGITO INVALIDO")	
	}

}
