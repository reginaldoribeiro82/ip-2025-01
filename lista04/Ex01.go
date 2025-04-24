package main

import "fmt"

func main (){

Arr01 := [10]int {}
var contador int

for i := 0; i < 10; i++ {
	fmt.Println("Digite um número inteiro")
	fmt.Scan(&Arr01[i])
	}

fmt.Println("Os seguintes números são maiores que 50 :")
for j := 0; j < 10 ; j++ {
	if Arr01[j] >= 50 {
	fmt.Print(" ",Arr01[j])
	contador = contador + 1
	}
}
	fmt.Println("")
	fmt.Print("O total de números que são maiores que 50 são ", contador)

}