package models

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type CalcRequest struct {
	Op           string  `json:"op,omitempty"`
	FirstNumber  *float64 `json:"firstNumber"`
	SecondNumber *float64 `json:"secondNumber"`
}

func (r *CalcRequest) Validate() error {
	if r.Op == "" && r.FirstNumber == nil && r.SecondNumber == nil {
		return fmt.Errorf("request body is empty or malformed")
	}

	if r.Op == "" {
		return errParamIsRequired("operator", "string")
	}

	if r.FirstNumber == nil {
		return errParamIsRequired("firstNumber", "float64")
	}

	if r.SecondNumber == nil {
		return errParamIsRequired("secondNumber", "float64")
	}
	return nil
}