package main

import "fmt"

func main() {
	var x interface{} = 10
	var y interface{} = "hello world "
	showType(x)
	showType(y)

}

func showType(t interface{}) {
	fmt.Printf("o tipo da variavel e %T e o valor e %v\n", t, t)
}
