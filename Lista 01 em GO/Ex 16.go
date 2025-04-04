package main

import "fmt"

func main() {
	var sal, sal_reaj float64

	fmt.Println("Digite o salario")
	fmt.Scan (&sal)
 
 
	if sal <= 300 {
	  sal_reaj = sal * 1.5
	  fmt.Printf("Salario com reajuste = %.2f", sal_reaj)
	} else {
	  sal_reaj = sal * 1.3
	  fmt.Printf("Salario com reajuste = %.2f", sal_reaj)
	}
}