package main

import "fmt"

func main() {

	Arr01 := [10]int{}
	contimpar := 0
	var Soma int
	var Arrpar []int
	var Arrimpar []int
	for i := 0; i < len(Arr01); i++ {
		fmt.Println("Digite um número inteiro")
		fmt.Scan(&Arr01[i])

	}

	for i := 0; i < len(Arr01); i++ {

		if Arr01[i]%2 == 0 {
			Arrpar = append(Arrpar, Arr01[i])
			Soma = Soma + Arr01[i]

		} else {
			Arrimpar = append(Arrimpar, Arr01[i])
			contimpar = contimpar + 1
		}
	}
	fmt.Print("Os números impares são: \n")
	fmt.Println(Arrpar)
	fmt.Printf("A soma dos números pares é %d", Soma)
	fmt.Println("")
	fmt.Println("Os números impares são :")
	fmt.Println(Arrimpar)
	fmt.Printf("O total de números impares é %d", contimpar)

}
