package validations

import (
	"github.com/binadel/vali/constraints"
	"github.com/binadel/vali/validations/results"
)

type IpValidation struct {
	stringValidation StringValidation
	version          int
}

func (v IpValidation) Version4() IpValidation {
	v.version = 4
	return v
}

func (v IpValidation) Version6() IpValidation {
	v.version = 6
	return v
}

// Validate applies the validations constraints to the field value and returns the result.
func (v IpValidation) Validate(value string) results.IpResult {
	var result results.IpResult
	result.StringResult = v.stringValidation.Validate(value)

	if len(result.Errors) == 0 {
		if ip, err := constraints.ParseIP(value, v.version); err != nil {
			result.Errors = append(result.Errors, err)
		} else {
			result.IP = ip
		}
	}

	return result
}
