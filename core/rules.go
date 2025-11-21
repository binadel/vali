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
)
