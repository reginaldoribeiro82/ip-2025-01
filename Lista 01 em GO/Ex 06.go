package main

import "fmt"

func main() {

	var fahr float64
	var celsius float64
	var n_temp int

	fmt.Println("Digite o número de temperaturas a serem convertidas")
	fmt.Scan(&n_temp)

	for i := 1; i <= n_temp; i++ {

		fmt.Println("Digite o valor em graus Fahrenheit")
		fmt.Scan(&fahr)

		celsius = 5 * (fahr - 32) / 9

		fmt.Printf("%.2f FAHRENHEIT EQUIVALE A %.2f CELSIUS\n", fahr, celsius)

	}

}
