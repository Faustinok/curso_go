package main

func main() {
	// array : tabamnho fixo e pode ser porcorrido
	var meu_array [3]int
	meu_array[0] = 1
	meu_array[1] = 2
	meu_array[2] = 13

	println(len(meu_array))
	println(meu_array[len(meu_array)-1])

	for i := 0; i < len(meu_array); i++ {
		println(meu_array[i])
	}
}
