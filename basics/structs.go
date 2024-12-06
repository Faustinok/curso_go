package main

import "fmt"

func main() {
	gabriel := Cliente{
		Nome:  "gabriel",
		Idade: 28,
		Ativo: true,
	}
	fmt.Println(gabriel.Nome)
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}
