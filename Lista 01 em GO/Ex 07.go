package main

import "fmt"

func main() {

	var fahr float64
	var celsius float64
	var chuva float64
	var milimetro float64

	fmt.Println("Digite o valor da temperatura em graus Fahrenheit")
	fmt.Scan(&fahr)
	fmt.Println("Digite o valor da quantidade de chuva em polegadas")
	fmt.Scan(&chuva)

	celsius = 5 * ((fahr - 32) / 9)
	milimetro = chuva * 25.4

	fmt.Printf("O VALOR EM CELSIUS = %.2f\n", celsius)
	fmt.Printf("A QUANTIDADE DE CHUVA E = %.2f", milimetro)

}
