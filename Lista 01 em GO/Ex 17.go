package main

import "fmt"

func main() {

	var num1, num2, seq int64
	var rest int64

	fmt.Println("Digite um número")
	fmt.Scan(&num1)
	fmt.Println("Digite um outro número")
	fmt.Scan(&num2)

	seq = num1 - 2
	rest = num1 % 2

	if rest == 0 {
		for i := 0; i < int(num2); i++ {
			seq = seq + 2
			fmt.Print(" ", seq)
		}
	} else {
		fmt.Println("O primeiro número não é par")
	}

}
