package main

import "fmt"

func main() {
	salarios := map[string]int{"gabriel": 6450.00, "guilherme": 12000}
	//fmt.Println(salarios["gabriel"])
	//fmt.Println(len(salarios))
	salarios["matheus"] = 20000
	fmt.Println(salarios)
	delete(salarios, "matheus")

	//sal := make(map[string]int)
	//sal1 := map[string]int{}
	for nome, salario := range salarios {
		fmt.Printf("o salario de %s e de %d\n", nome, salario)
	}
}
