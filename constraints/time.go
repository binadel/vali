package constraints

import (
	"time"

	"github.com/binadel/vali/core"
	"github.com/binadel/vali/errors"
)

var dateError = &errors.BasicError{
	Code:    core.RuleDate,
	Message: "value must be a valid date",
}

var timeError = &errors.BasicError{
	Code:    core.RuleTime,
	Message: "value must be a valid time",
}

var dateTimeError = &errors.BasicError{
	Code:    core.RuleDateTime,
	Message: "value must be a valid date time",
}

var timeLayouts = []string{
	"15:04:05Z07:00",
	"15:04:05.999999999Z07:00",
	"15:04:05",
	"15:04:05.999999999",
	"15:04",
}

func ParseDate(value string) (time.Time, error) {
	if date, err := time.Parse(time.DateOnly, value); err == nil {
		return date, nil
	}
	return time.Time{}, dateError
}

func ParseTime(value string) (time.Time, error) {
	// Try all valid RFC3339 time layouts
	for _, layout := range timeLayouts {
		if t, err := time.Parse(layout, value); err == nil {
			return t, nil
		}
	}
	return time.Time{}, timeError
}

func ParseDateTime(value string) (time.Time, core.Error) {
	// JSON Schema uses RFC3339
	if dateTime, err := time.Parse(time.RFC3339Nano, value); err == nil {
		return dateTime, nil
	}
	return time.Time{}, dateTimeError
}
