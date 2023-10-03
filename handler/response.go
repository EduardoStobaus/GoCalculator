package handler

type CalcResponse struct {
	Error  string  `json:"error"`
	Result float64 `json:"result"`
}
