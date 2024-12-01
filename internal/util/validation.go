package util

import "regexp"

var webPattern *regexp.Regexp = regexp.MustCompile(`^[a-zA-Z\d_]+$`)

func ValidateWeb(s string) bool {
	return webPattern.MatchString(s)
}
