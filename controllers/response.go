package controllers

type Response struct {
	Erro string  `json:"error,omitempty"`
	Resultado float64 `json:"result"`
}
