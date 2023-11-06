package models

import (
	"fmt"
)

type CalcRequest struct {
	Op           string  `json:"op,omitempty"`
	FirstNumber  float64 `json:"firstNumber"`
	SecondNumber float64 `json:"secondNumber"` 
}

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

func (r *CalcRequest) Validate() error {
	if r.Op == "" && r.FirstNumber == 0 && r.SecondNumber == 0 {
		return fmt.Errorf("request body is empty or malformed")
	}

	if r.Op == "" {
		return errParamIsRequired("operator", "string")
	}
	return nil
}
