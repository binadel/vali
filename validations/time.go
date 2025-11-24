package validations

import (
	"github.com/binadel/vali/constraints"
	"github.com/binadel/vali/validations/results"
)

type TimeValidation struct {
	stringValidation StringValidation
	onlyTime         bool
	onlyDate         bool
}

// Validate applies the validations constraints to the field value and returns the result.
func (v TimeValidation) Validate(value string) results.TimeResult {
	var result results.TimeResult
	result.StringResult = v.stringValidation.Validate(value)

	if len(result.Errors) == 0 {
		if v.onlyTime {
			if t, err := constraints.ParseTime(value); err != nil {
				result.Errors = append(result.Errors, err)
			} else {
				result.Time = t
			}
		} else if v.onlyDate {
			if t, err := constraints.ParseDate(value); err != nil {
				result.Errors = append(result.Errors, err)
			} else {
				result.Time = t
			}
		} else {
			if t, err := constraints.ParseDateTime(value); err != nil {
				result.Errors = append(result.Errors, err)
			} else {
				result.Time = t
			}
		}
	}

	return result
}
