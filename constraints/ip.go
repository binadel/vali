package constraints

import (
	"net"

	"github.com/binadel/vali/core"
	"github.com/binadel/vali/errors"
)

var ipError = &errors.BasicError{
	Code:    core.RuleIP,
	Message: "value must be a valid IP address",
}

var ipv4Error = &errors.BasicError{
	Code:    core.RuleIPv4,
	Message: "value must be a valid IPv4 address",
}

var ipv6Error = &errors.BasicError{
	Code:    core.RuleIPv6,
	Message: "value must be a valid IPv6 address",
}

func ParseIP(value string, version int) (net.IP, core.Error) {
	if ip := net.ParseIP(value); ip != nil {
		if version == 4 {
			if ipv4 := ip.To4(); ipv4 != nil {
				return ipv4, nil
			} else {
				return nil, ipv4Error
			}
		}
		if version == 6 {
			if ipv4 := ip.To4(); ipv4 == nil {
				return ip, nil
			} else {
				return nil, ipv6Error
			}
		}
		return ip, nil
	}
	return nil, ipError
}
