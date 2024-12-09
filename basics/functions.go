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

type Endereco struct {
	Logradouro string
	Numero     string
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (c Cliente) Desativar() {
	c.Ativo = false
}
