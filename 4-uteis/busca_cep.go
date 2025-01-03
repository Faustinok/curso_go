package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"dd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, url := range os.Args[1:] {
		req, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error occured: %s\n", err)

		}
		defer req.Body.Close()
		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error occured: %s\n", err)
		}
		var data ViaCep
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error occured: %s\n", err)
		}
		file, err := os.Create("cidade.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error occured: %s\n", err)
		}
		defer file.Close()
		_, err = file.WriteString(fmt.Sprintf("Cep: %s, Localidade: %s, Uf: %s", data.Cep, data.Localidade, data.Uf))

	}
}
