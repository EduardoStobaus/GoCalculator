package models

import (
	"fmt"
)

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type CalcRequest struct {
	Op           string  `json:"op,omitempty"`
	FirstNumber  float64 `json:"firstNumber"`
	SecondNumber float64 `json:"secondNumber"`
}

func (r *CalcRequest) Validate() error {
	if r.Op == "" {
		return fmt.Errorf("request body is empty or malformed")
	}

	if r.Op == "" {
		return errParamIsRequired("operator", "string")
	}
	return nil
}
