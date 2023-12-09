package bb

import (
	"regexp"
	"strings"

	"github.com/gobwas/glob"
)

var (
	CaseSensitive = false
)

func IsRegExp(pattern string) bool {
	_, err := regexp.Compile(pattern)
	return err == nil
}

func IsGlobPattern(pattern string) bool {
	_, err := glob.Compile(pattern)
	return err == nil
}

func StringMatchesGlobPattern(value, pattern string) (bool, error) {
	if !CaseSensitive {
		value = strings.ToLower(value)
		pattern = strings.ToLower(pattern)
	}
	g, err := glob.Compile(pattern)
	if err != nil {
		return false, err
	}
	return g.Match(value), nil
}

func StringMatchesRegExp(value, pattern string) (bool, error) {
	if !CaseSensitive {
		value = strings.ToLower(value)
		pattern = strings.ToLower(pattern)
	}
	r, err := regexp.Compile(pattern)
	if err != nil {
		return false, err
	}
	return r.MatchString(value), nil
}

func StringMatchesPattern(value, pattern string) (bool, error) {
	if pattern == "" {
		return true, nil
	}
	if IsRegExp(pattern) {
		return StringMatchesRegExp(value, pattern)
	}
	if IsGlobPattern(pattern) {
		return StringMatchesGlobPattern(value, pattern)
	}
	return false, nil
}

func StringMatchesAnyPattern(value string, patterns []string) (bool, error) {
	for _, pattern := range patterns {
		matches, err := StringMatchesPattern(value, pattern)
		if err != nil {
			return false, err
		}
		if matches {
			return true, nil
		}
	}
	return false, nil
}

func AnyStringMatchesPattern(values []string, pattern string) (bool, error) {
	for _, value := range values {
		matches, err := StringMatchesPattern(value, pattern)
		if err != nil {
			return false, err
		}
		if matches {
			return true, nil
		}
	}
	return false, nil
}

func AnyStringMatchesAnyPattern(values, patterns []string) (bool, error) {
	for _, value := range values {
		matches, err := StringMatchesAnyPattern(value, patterns)
		if err != nil {
			return false, err
		}
		if matches {
			return true, nil
		}
	}
	return false, nil
}
