package main

import "fmt"

func main() {
	
var sal_min, gasto_kwh, kwh_sal, gasto_reais, gasto_reais_desc float64

fmt.Println("Digite o valor do salário minimo")
fmt.Scan(&sal_min)
fmt.Println("Digite o consumo de kwh da residência")
fmt.Scan(&gasto_kwh)

kwh_sal = (sal_min * 0.007)
gasto_reais = (kwh_sal * gasto_kwh)
gasto_reais_desc = (gasto_reais * (0.9))

fmt.Printf("Custo por kW: R$ %.2f", kwh_sal)
fmt.Println("")
fmt.Printf("Custo do consumo: R$ %.2f", gasto_reais)
fmt.Println("")
fmt.Printf("Custo com desconto: R$ %.2f", gasto_reais_desc)
}