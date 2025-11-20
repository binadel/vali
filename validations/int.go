package validations

import (
	"github.com/binadel/vali/constraints"
	"github.com/binadel/vali/core"
	"github.com/binadel/vali/validations/results"
)

// IntValidation defines the validation constraints for an integer field.
type IntValidation struct {
	Path        core.FieldPath
	constraints []core.IntValidator
}

// Minimum adds a minimum constraint to the field.
func (v IntValidation) Minimum(minimum int64) IntValidation {
	v.constraints = append(v.constraints, constraints.MinimumInt(minimum))
	return v
}

// Maximum adds a maximum constraint to the field.
func (v IntValidation) Maximum(maximum int64) IntValidation {
	v.constraints = append(v.constraints, constraints.MaximumInt(maximum))
	return v
}

// ExclusiveMinimum adds an exclusive minimum constraint to the field.
func (v IntValidation) ExclusiveMinimum(exclusiveMinimum int64) IntValidation {
	v.constraints = append(v.constraints, constraints.ExclusiveMinimumInt(exclusiveMinimum))
	return v
}

// ExclusiveMaximum adds an exclusive maximum constraint to the field.
func (v IntValidation) ExclusiveMaximum(exclusiveMaximum int64) IntValidation {
	v.constraints = append(v.constraints, constraints.ExclusiveMaximumInt(exclusiveMaximum))
	return v
}

// MultipleOf adds a multipleOf constraint to the field.
func (v IntValidation) MultipleOf(multiplier int64) IntValidation {
	v.constraints = append(v.constraints, constraints.MultipleOfInt(multiplier))
	return v
}

// Validate applies the validations constraints to the field value and returns the result.
func (v IntValidation) Validate(value int64) results.IntResult {
	var errors []core.Error
	for _, constraint := range v.constraints {
		if err := constraint(value); err != nil {
			errors = append(errors, err)
		}
	}
	return results.IntResult{
		Path:   v.Path,
		Value:  value,
		Errors: errors,
	}
}
