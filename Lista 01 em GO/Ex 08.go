package main

import "fmt"

func main() {

	var raio, altura, custo, area_base, area_total, area_lateral float64
	const pi float64 = 3.14159

	fmt.Println("Digite o raio da lata em metros")
	fmt.Scan(&raio)
	fmt.Println("Digite a altura da lata em metros")
	fmt.Scan(&altura)

	area_base = pi * (raio * raio)
	area_lateral = 2 * pi * raio * altura
	area_total = (2 * area_base) + area_lateral
	custo = area_total * 100

	fmt.Printf("O VALOR DO CUSTO E = %.2f", custo)

}
