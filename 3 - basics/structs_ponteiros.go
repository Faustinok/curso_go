package main

import "fmt"

type Cliente struct {
	nome string
}

func (c *Cliente) andou() {
	c.nome = "gabriel caralho"
	fmt.Printf("o cliente %v andou\n ", c.nome)
}
func main() {
	gabriel := Cliente{nome: "gabriel"}
	gabriel.andou()
	fmt.Printf("o nome dele e %v", gabriel.nome)
}
