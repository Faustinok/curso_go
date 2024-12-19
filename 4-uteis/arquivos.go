package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}
	//tamanho, err := f.Write([]byte("hellor world")) salvar bytes
	tamanho, err := f.WriteString("hello world")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso tamanho = %d\n", tamanho)
	// leitura
	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(arquivo))
	// leitura pco a pco
	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 10)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		println(string(buffer[:n]))
	}
	f.Close()
	arquivo2.Close()
	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}
}
