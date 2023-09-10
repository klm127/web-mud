package re

import "regexp"

var HasPeriod regexp.Regexp

func init() {
	HasPeriod = *regexp.MustCompile(`.*\..*`)
}
