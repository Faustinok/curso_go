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

type Endereco2 struct {
	Logradouro string
	Numero     string
	Cidade     string
	Estado     string
}

type Cliente2 struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco2
}
