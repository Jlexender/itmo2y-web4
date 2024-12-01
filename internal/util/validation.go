package util

import "regexp"

var webPattern = regexp.MustCompile(`^[a-zA-Z\d_]+$`)

func IsCredential(s string) bool {
	return webPattern.MatchString(s)
}
