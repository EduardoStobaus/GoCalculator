package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/EduardoStobaus/GoCalculator/schemas"
)

func Calculator(w http.ResponseWriter, r *http.Request) {
	req := &schemas.CalcRequest{}
	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		fmt.Println("Erro no Body Request: ", err)
	}

	// Operação
	resp := Response{}
	switch req.Op {
	case "+":
		resp.Result = req.FirstNumber + req.SecondNumber
	case "-":
		resp.Result = req.FirstNumber - req.SecondNumber
	case "*":
		resp.Result = req.FirstNumber * req.SecondNumber
	case "/":
		if req.SecondNumber == 0.0 {
			resp.Error = "Divisão por 0 é impossível"
		} else {
			resp.Result = req.FirstNumber / req.SecondNumber
		}
	default:
		resp.Error = fmt.Sprintf("Operação desconhecida: %s! Tente novamente.", req.Op)
	}

	w.Header().Set("Content-Type", "application/json")
	if resp.Error != "" {
		w.WriteHeader(http.StatusBadRequest)
	}

	enc := json.NewEncoder(w)
	if err := enc.Encode(resp); err != nil {
		log.Printf("Não deu pra codificar %v - %s", resp, err)
	}
}
