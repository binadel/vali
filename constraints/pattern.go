package constraints

import (
	"regexp"
	"sync"

	"github.com/binadel/vali/core"
	"github.com/binadel/vali/errors"
)

var patternCache sync.Map

func getRegex(pattern string) *regexp.Regexp {
	if regex, ok := patternCache.Load(pattern); ok {
		return regex.(*regexp.Regexp)
	}

	regex := regexp.MustCompile(pattern)
	actual, _ := patternCache.LoadOrStore(pattern, regex)
	return actual.(*regexp.Regexp)
}

var patternError = &errors.BasicError{
	Code:    core.RulePattern,
	Message: "value does not match the required pattern",
}

var regexError = &errors.BasicError{
	Code:    core.RuleRegex,
	Message: "value must a valid regular expression",
}

// PatternString applies regex pattern constraint to a string.
func PatternString(pattern string) core.StringValidator {
	regex := getRegex(pattern)
	return func(value string) core.Error {
		if !regex.MatchString(value) {
			return patternError
		}
		return nil
	}
}

// ParseRegex parses the given string as a regular expression.
func ParseRegex(value string) (*regexp.Regexp, core.Error) {
	regex, err := regexp.Compile(value)
	if err != nil {
		return nil, regexError
	}
	return regex, nil
}
