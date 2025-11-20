package core

import "github.com/mailru/easyjson"

// Error represents a validation error with two forms of representation:
//  1. the standard Go error string, and
//  2. a structured JSON form via easyjson.
//
// Implementations can be used as normal errors while also supporting
// efficient, zero-allocation JSON serialization.
type Error interface {
	error
	easyjson.Marshaler
}
