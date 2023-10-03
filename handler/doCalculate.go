package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func doCalculateHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	dec := json.NewDecoder(r.Body) //Criando novo Decoder
	req := &CalcRequest{}          //Criando request vazio

	if err := dec.Decode(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Operação
	resp := &CalcResponse{}
	switch req.Op {
	case "+":
		resp.Result = req.Left + req.Right
	case "-":
		resp.Result = req.Left - req.Right
	case "*":
		resp.Result = req.Left * req.Right
	case "/":
		if req.Right == 0.0 {
			resp.Error = "divisão por 0"
		} else {
			resp.Result = req.Left / req.Right
		}
	default:
		resp.Error = fmt.Sprintf("Operação desconhecida: %s", req.Op)
	}

	w.Header().Set("Content-Type", "application/json")
	if resp.Error != "" {
		w.WriteHeader(http.StatusBadRequest)
	}

	enc := json.NewEncoder(w)
	if err := enc.Encode(resp); err != nil {
		log.Printf("Não dá pra fazer encode %v - %s", resp, err)
	}

}
