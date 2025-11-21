package validations

import (
	"github.com/binadel/vali/constraints"
	"github.com/binadel/vali/core"
	"github.com/binadel/vali/validations/results"
)

// StringValidation defines the validation constraints for an integer field.
type StringValidation struct {
	Path        core.FieldPath
	constraints []core.StringValidator
}

// Length adds a length constraint to the field.
func (v StringValidation) Length(expectedLength int) StringValidation {
	v.constraints = append(v.constraints, constraints.LengthString(expectedLength))
	return v
}

// MinLength adds a min length constraint to the field.
func (v StringValidation) MinLength(minLength int) StringValidation {
	v.constraints = append(v.constraints, constraints.MinLengthString(minLength))
	return v
}

// MaxLength adds a max length constraint to the field.
func (v StringValidation) MaxLength(maxLength int) StringValidation {
	v.constraints = append(v.constraints, constraints.MaxLengthString(maxLength))
	return v
}

// Pattern adds a regex pattern constraint to the field.
func (v StringValidation) Pattern(pattern string) StringValidation {
	v.constraints = append(v.constraints, constraints.PatternString(pattern))
	return v
}

// Validate applies the validations constraints to the field value and returns the result.
func (v StringValidation) Validate(value string) results.StringResult {
	var errors []core.Error
	for _, constraint := range v.constraints {
		if err := constraint(value); err != nil {
			errors = append(errors, err)
		}
	}
	return results.StringResult{
		Path:   v.Path,
		Value:  value,
		Errors: errors,
	}
}
