package results

import (
	"github.com/binadel/vali/core"
	"github.com/mailru/easyjson/jwriter"
)

// IntResult represents the validation result for an integer field.
type IntResult struct {
	Path   core.FieldPath
	Value  int64
	Errors []core.Error
}

// IsValid returns whether the field has passed validation test.
func (r IntResult) IsValid() bool {
	return len(r.Errors) == 0
}

// MarshalEasyJSON writes the json representation to the output.
func (r IntResult) MarshalEasyJSON(w *jwriter.Writer) {
	w.RawByte('{')

	w.RawString(keyPath)
	r.Path.MarshalEasyJSON(w)

	w.RawString(keyValue)
	w.Int64(r.Value)

	w.RawString(keyErrors)
	w.RawByte('[')
	for i, err := range r.Errors {
		if i > 0 {
			w.RawByte(',')
		}
		err.MarshalEasyJSON(w)
	}
	w.RawByte(']')

	w.RawByte('}')
}
