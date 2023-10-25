package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/EduardoStobaus/GoCalculator/schemas"
)

func Calculadora(w http.ResponseWriter, r *http.Request) {
	req := &schemas.CalcRequest{}
	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		fmt.Println("Erro no Body Request: ", err)
	}

	// Operação
	resp := Response{}
	switch req.Operador {
	case "+":
		resp.Resultado = req.PrimeiroNumero + req.SegundoNumero
	case "-":
		resp.Resultado = req.PrimeiroNumero - req.SegundoNumero
	case "*":
		resp.Resultado = req.PrimeiroNumero * req.SegundoNumero
	case "/":
		if req.SegundoNumero == 0.0 {
			resp.Erro = "Bad Request: Divisão por 0 não é válido"
		} else {
			resp.Resultado = req.PrimeiroNumero / req.SegundoNumero
		}
	default:
		resp.Erro = fmt.Sprintf("Bad Request: Operador não definido corretamente: %s", req.Operador)
	}

	w.Header().Set("Content-Type", "application/json")
	if resp.Erro != "" {
		w.WriteHeader(http.StatusBadRequest)
	}

	enc := json.NewEncoder(w)
	if err := enc.Encode(resp); err != nil {
		log.Printf("Não foi possível codificar")
	}

}
