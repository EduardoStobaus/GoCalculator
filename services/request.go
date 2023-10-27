package services

type CalcRequest struct {
	Op       string  `json:"op,omitempty"`
	FirstNumber float64 `json:"firstNumber"`
	SecondNumber  float64 `json:"secondNumber"`
}
