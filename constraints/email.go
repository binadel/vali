package constraints

import (
	"net/mail"

	"github.com/binadel/vali/core"
	"github.com/binadel/vali/errors"
)

var emailError = &errors.BasicError{
	Code:    core.RuleEmail,
	Message: "value must be a valid email",
}

// ParseEmail parses the given string as an email address.
func ParseEmail(value string) (*mail.Address, core.Error) {
	email, err := mail.ParseAddress(value)
	if err != nil {
		return nil, emailError
	}
	return email, nil
}
