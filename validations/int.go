package validations

import (
	"github.com/binadel/vali/constraints"
	"github.com/binadel/vali/core"
	"github.com/mailru/easyjson/jwriter"
)

const (
	keyPath   = "\"path\":"
	keyValue  = ",\"value\":"
	keyErrors = ",\"errors\":"
)

type IntResult struct {
	path   core.FieldPath
	value  int64
	errors []core.Error
}

func (r IntResult) Path() core.FieldPath {
	return r.path
}

func (r IntResult) Value() int64 {
	return r.value
}

func (r IntResult) Errors() []core.Error {
	return r.errors
}

func (r IntResult) MarshalEasyJSON(w *jwriter.Writer) {
	w.RawByte('{')

	w.RawString(keyPath)
	r.path.MarshalEasyJSON(w)

	w.RawString(keyValue)
	w.Int64(r.value)

	w.RawString(keyErrors)
	w.RawByte('[')
	for i, err := range r.errors {
		if i > 0 {
			w.RawByte(',')
		}
		err.MarshalEasyJSON(w)
	}
	w.RawByte(']')

	w.RawByte('}')
}

type IntValidation struct {
	Path        core.FieldPath
	constraints []core.IntValidator
}

func (v IntValidation) Min(min int64) IntValidation {
	v.constraints = append(v.constraints, constraints.MinInt(min))
	return v
}

func (v IntValidation) Validate(value int64) IntResult {
	var errors []core.Error
	for _, constraint := range v.constraints {
		if err := constraint(value); err != nil {
			errors = append(errors, err)
		}
	}
	return IntResult{
		path:   v.Path,
		value:  value,
		errors: errors,
	}
}
