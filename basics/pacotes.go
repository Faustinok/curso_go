package main

import (
	matematica "curso/matematica"
	"fmt"
)

func main() {
	soma := matematica.Soma(10, 20)
	fmt.Printf("resultado = %v", soma)
}
