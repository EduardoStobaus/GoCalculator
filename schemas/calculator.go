package schemas

type CalcRequest struct {
	Op           string  `json:"op"`
	FirstNumber  float64 `json:"first_number"`
	SecondNumber float64 `json:"second_number"`
}
