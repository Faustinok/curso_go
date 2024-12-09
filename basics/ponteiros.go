package main

func main() {
	a := 10
	println(&a) //-> endereco na memoria
	println(a)
	var ponteiro *int = &a // o * indica que e um ponteiro
	*ponteiro = 20
	println(a)
}
