package main

import "fmt"

func main() {

	var horas, horas_fx1, horas_fx2, total float64

	fmt.Println("Digite a quantidade de horas que a charrete foi usada")
	fmt.Scan(&horas)

	if horas < 2 {
		horas_fx1 = 5 * horas
		fmt.Printf("O VALOR A PAGAR É = %f \n", horas_fx1)
	} else {
		horas_fx2 = 10 * ((horas - 2) / 3)
		total = horas_fx2 + 10
		fmt.Printf("O VALOR A PAGAR É = %.2f \n", total)
	}

}
