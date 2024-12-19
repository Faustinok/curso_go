package main

import (
	"encoding/json"
)

type Conta struct {
	Numero int
	Saldo  int
}

func main() {
	conta := Conta{Numero: 10, Saldo: 20}
	res, err := json.Marshal(conta)
	if err != nil {
		print(err)
	}
	println(string(res))
	//encoder := json.NewEncoder(os.Stdout)
	//encoder.Encode(conta)

	// criar um obj baseado em um json

	var contaX Conta
	jsonpuro := []byte(`{"numero":30, "saldo":40}`)
	err = json.Unmarshal(jsonpuro, &contaX)
	if err != nil {
		print(err)
	}
	println(contaX.Numero, conta.Saldo)
}
