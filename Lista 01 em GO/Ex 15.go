package main

import "fmt"

func main() {

	var num float64
	var quadrado float64

	fmt.Println("Digite um número inteiro (N) 5 < N < 2000")
	fmt.Scan(&num)

	if num > 5 && num < 2000 {
		quadrado = num * num
		fmt.Println(num, "^", num, "=", quadrado)
	}

}
