package templateutils

import (
	"regexp"
	"strings"
)

var (
	intTypeRegex     = regexp.MustCompile(`^[u]?int`)
	floatTypeRegex   = regexp.MustCompile(`^float`)
	complexTypeRegex = regexp.MustCompile(`^complex`)
)

// ZeroValue will return input type name's default zero value.
func ZeroValue(in string) string {
	switch in {
	case "bool":
		return "false"
	case "string":
		return `""`
	}
	// Handle all integer, float and complex via return 0.
	if intTypeRegex.MatchString(in) || floatTypeRegex.MatchString(in) || complexTypeRegex.MatchString(in) {
		return "0"
	}

	// If type starts with a "*", it must be a pointer, return nil directly.
	if strings.HasPrefix(in, "*") {
		return "nil"
	}

	// If type starts with "[]", it must a slice, return nil directly.
	if strings.HasPrefix(in, "[]") {
		return "nil"
	}

	return in + "{}"
}
