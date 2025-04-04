package main

import "fmt"

func main() {
	var n_teste, total, n_popular, n_geral, n_arquibancada, n_cadeiras float64
	var renda_jogo float64
	var i float64

	fmt.Println("Digite o numero de testes")
	fmt.Scan(&n_teste)

	for i = 1; i <= n_teste; i++ {
		fmt.Println("Digite o número de pessoas que comparam ingressos")
		fmt.Scan(&total)

		fmt.Println("Digite a percentagem de pessoas que comparam ingresos na categoria (Popular)")
		fmt.Scan(&n_popular)

		fmt.Println("Digite a percentagem de pessoas que comparam ingresos na categoria (Geral)")
		fmt.Scan(&n_geral)

		fmt.Println("Digite a percentagem de pessoas que comparam ingresos na categoria (Arquibancada)")
		fmt.Scan(&n_arquibancada)

		fmt.Println("Digite a percentagem de pessoas que comparam ingresos na categoria (Cadeiras)")
		fmt.Scan(&n_cadeiras)

		renda_jogo = (total * (n_popular / 100) * 1) + (total * (n_geral / 100) * 5) + (total * (n_arquibancada / 100) * 10) + (total * (n_cadeiras / 100) * 20)
		fmt.Printf("A renda do jogo Nº %.1f = %.2f \n", i, renda_jogo)

	}

}
