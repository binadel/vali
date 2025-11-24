package constraints

import (
	"net/url"

	"github.com/binadel/vali/core"
	"github.com/binadel/vali/errors"
	"github.com/google/uuid"
)

var uuidError = &errors.BasicError{
	Code:    core.RuleUuid,
	Message: "value must be a valid UUID",
}

var uriError = &errors.BasicError{
	Code:    core.RuleUri,
	Message: "value must be a valid URI",
}

// ParseUuid parses the given string as a UUID.
func ParseUuid(value string) (uuid.UUID, core.Error) {
	if id, err := uuid.Parse(value); err == nil {
		return id, nil
	}
	return uuid.Nil, uuidError
}

// ParseUri parses the given string as a URI.
func ParseUri(value string) (*url.URL, core.Error) {
	if uri, err := url.Parse(value); err == nil {
		return uri, nil
	}
	return nil, uriError
}
