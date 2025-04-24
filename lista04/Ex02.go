package main

import "fmt"

func main (){

Arr01 := [10]int {}
Arr02 := [5]int {}
var Arr03 [] int
var Arr04 [] int

fmt.Println("Primeiro Vetor")
for i := 0; i < len(Arr01); i++ {
	fmt.Println("Digite um número inteiro")
	fmt.Scan(&Arr01[i])
}

fmt.Println("Segundo Vetor")
for i := 0; i < len(Arr02); i++ {
	fmt.Println("Digite um número inteiro")
	fmt.Scan(&Arr02[i])
}

fmt.Println("Primeiro Vetor Resultante")

for i := 0; i < len(Arr01) ; i++ {
	if Arr01[i] % 2 ==0 {
		Arr03 = append(Arr03,Arr01[i] + Arr02[0] + Arr02[1] + Arr02[2] + Arr02[3] + Arr02[4])
	}
}

fmt.Print(Arr03)
fmt.Println("")

fmt.Println("Segundo Vetor Resultante")
for i := 0; i < len(Arr01) ; i++ {
	if Arr01[i] % 2 !=0 {
		Arr04 = append(Arr04,Arr01[i] + Arr02[0] + Arr02[1] + Arr02[2] + Arr02[3] + Arr02[4])
	}
}

fmt.Print(Arr04)
}
