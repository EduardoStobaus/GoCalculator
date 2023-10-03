package handler

type CalcRequest struct {
	Op    string  `json:"op"`
	Left  float64 `json:"left"`
	Right float64 `json:"right"`
}
