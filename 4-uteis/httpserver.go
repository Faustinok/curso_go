package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type ViaCeps struct {
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
	http.HandleFunc("/", BuscacepHandles)
	http.ListenAndServe(":8080", nil)

}
func BuscacepHandles(w http.ResponseWriter, r *http.Request) {
	cepParam := r.URL.Query().Get("cep")
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	cep, error := Buscacep(cepParam)
	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello world "))
	json.NewEncoder(w).Encode(cep)

}
func Buscacep(cep string) (*ViaCeps, error) {
	resp, error := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if error != nil {
		return nil, error
	}
	defer resp.Body.Close()
	body, error := io.ReadAll(resp.Body)
	if error != nil {
		return nil, error
	}
	var c ViaCeps
	error = json.Unmarshal(body, &c)
	if error != nil {
		return nil, error
	}
	return &c, nil
}
