package constraints

import (
	"github.com/binadel/vali/core"
	"github.com/binadel/vali/errors"
)

var maxError = errors.BasicError{
	Code:    core.RuleMax,
	Message: "field value is greater than allowed maximum",
}

func MaxInt(max int64) core.IntValidator {
	return func(value int64) core.Error {
		if value > max {
			return &errors.IntParamError{
				BasicError: maxError,
				ParamKey:   errors.ParamKeyLimit,
				ParamValue: max,
			}
		}
		return nil
	}
}
