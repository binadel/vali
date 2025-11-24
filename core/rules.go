package core

// Rule is the name for the field constraint.
type Rule string

const (
	RuleMinimum          Rule = "MINIMUM"
	RuleMaximum          Rule = "MAXIMUM"
	RuleExclusiveMinimum Rule = "EXCLUSIVE_MINIMUM"
	RuleExclusiveMaximum Rule = "EXCLUSIVE_MAXIMUM"
	RuleMultipleOf       Rule = "MULTIPLE_OF"
	RuleLength           Rule = "LENGTH"
	RuleMinLength        Rule = "MIN_LENGTH"
	RuleMaxLength        Rule = "MAX_LENGTH"
	RulePattern          Rule = "PATTERN"
	RuleDateTime         Rule = "DATE_TIME"
	RuleDate             Rule = "DATE"
	RuleTime             Rule = "TIME"
	RuleDuration         Rule = "DURATION"
	RuleEmail            Rule = "EMAIL"
	RuleIP               Rule = "IP"
	RuleIPv4             Rule = "IP_V4"
	RuleIPv6             Rule = "IP_V6"
	RuleUri              Rule = "URI"
	RuleUuid             Rule = "UUID"
	RuleRegex            Rule = "REGEX"
)
