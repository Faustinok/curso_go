package main

import "fmt"

func main() {
	//slices possuem ponteiros, tamanho e capacidade

	s := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Printf("len= %d cap=%d %v\n", len(s), cap(s), s)
	fmt.Printf("len= %d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])
	fmt.Printf("len= %d cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4])
	fmt.Printf("len= %d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:])
	s = append(s, 110) // ele dobra o tamanho do slide
	fmt.Printf("len= %d cap=%d %v\n", len(s), cap(s), s)
}
