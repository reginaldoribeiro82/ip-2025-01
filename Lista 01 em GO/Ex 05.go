package main

import "fmt"

func main() {

	var g_residencial float64
	var g_comercial_fx1 float64
	var g_comercial_fx2 float64
	var metro_c float64
	var g_industrial_fx1 float64
	var g_industrial_fx2 float64
	var metro_i float64
	var num float64
	var conta_cliente int64
	var tipo_c string

	fmt.Println("Digite o número da conta")
	fmt.Scan(&conta_cliente)
	fmt.Println("Digite o consumo de agua por metros cúbicos")
	fmt.Scan(&num)
	fmt.Println("Digite o tipo de consumidor (C - Comercial, I - Industrial, R -  Residencial)")
	fmt.Scan(&tipo_c)

	if tipo_c == "C" {
		if num <= 80 {
			g_comercial_fx1 = 500
			fmt.Printf("CONTA = %d\n", conta_cliente)
			fmt.Printf("VALOR DA CONTA = %.2f\n", g_comercial_fx1)
		} else {
			metro_c = num - 80
			g_comercial_fx2 = (metro_c * 0.25) + 500
			fmt.Printf("CONTA = %d\n", conta_cliente)
			fmt.Printf("VALOR DA CONTA = %.2f\n", g_comercial_fx2)
		}

	} else {
		if tipo_c == "I" {
			if num <= 100 {
				g_industrial_fx1 = 800
				fmt.Printf("CONTA = %d\n", conta_cliente)
				fmt.Printf("VALOR DA CONTA = %.2f\n", g_industrial_fx1)
			} else {
				metro_i = num - 100
				g_industrial_fx2 = (metro_i * 0.04) + 800
				fmt.Printf("CONTA = %.d\n", conta_cliente)
				fmt.Printf("VALOR DA CONTA = %.2f\n", g_industrial_fx2)

			}

		} else {
			g_residencial = 5 * num
			fmt.Printf("CONTA = %d\n", conta_cliente)
			fmt.Printf("VALOR DA CONTA = %.2f \n", g_residencial)
		}

	}

}
