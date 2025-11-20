package core

import (
	"strings"

	"github.com/mailru/easyjson/jwriter"
)

// PathRepresentation determines how to represent a path, which can be:
// 1. in array form: ["user", "address", "street"]
// 2. joined by dots: "user.address.street"
// 3. joined by slashes: "user/address/street"
type PathRepresentation int

const (
	PathRepresentationArray PathRepresentation = iota
	PathRepresentationDotted
	PathRepresentationSlashed
)

// FieldPathRepresentation is a global config for path representation.
var FieldPathRepresentation = PathRepresentationArray

// FieldPath represent the path to the field.
type FieldPath struct {
	segments []string
	json     []byte
}

// Field creates a new field path.
func Field(path []string) FieldPath {
	var json string
	switch FieldPathRepresentation {
	case PathRepresentationArray:
		json = "[\"" + strings.Join(path, "\",\"") + "\"]"
	case PathRepresentationDotted:
		json = "\"" + strings.Join(path, ".") + "\""
	case PathRepresentationSlashed:
		json = "\"" + strings.Join(path, "/") + "\""
	default:
		panic("invalid FieldPath representation type")
	}
	return FieldPath{
		segments: path,
		json:     []byte(json),
	}
}

// Segments return the path segments.
func (p FieldPath) Segments() []string {
	return p.segments
}

// MarshalEasyJSON writes the json representation to the output.
func (p FieldPath) MarshalEasyJSON(w *jwriter.Writer) {
	w.Raw(p.json, nil)
}
