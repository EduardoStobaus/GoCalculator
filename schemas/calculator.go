package schemas

type CalcRequest struct {
	Operador       string  `json:"operador"`
	PrimeiroNumero float64 `json:"primeiro_numero"`
	SegundoNumero  float64 `json:"segundo_numero"`
}
