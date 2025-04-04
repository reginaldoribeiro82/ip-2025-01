package main

import "fmt"

func main() {

	var altura_pir, aresta_hex, vol_pir, area_base float64

	fmt.Println("Digite a altura da pirade")
	fmt.Scan(&altura_pir)
	fmt.Println("Digite a aresta da aresta_hex")
	fmt.Scan(&aresta_hex)

	area_base = 3 * (aresta_hex * aresta_hex) * 0.866
	vol_pir = (area_base * altura_pir) / 3

	fmt.Printf("O VOLUME DA PIRADE É %.1f METROS CUBICOS", vol_pir)

}
