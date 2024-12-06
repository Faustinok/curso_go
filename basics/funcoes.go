package main

func main() {
	//closure

	total_ := func() int {
		return somas(2, 5, 7, 6, 5, 4, 566, 5, 5, 6) * 2
	}()
	println(total_)
	// fim closure
	println(soma(5, 2))
	println(total_)
}
func soma(a int, b int) (int, bool) {
	maior := false
	if a > 3 {
		maior = true
	}
	return a + b, maior
}
func somas(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
