package main

func Soma[T int | float64](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}
func main() {
	mapa := map[string]int{"Gabriel": 1000, "pedro": 23}
	println(Soma(mapa))
}
