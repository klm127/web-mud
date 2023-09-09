package language

import (
	"errors"
	"strings"
)

var yes = []string{"yes", "y", "indeed", "ok", "affirmative"}
var nos = []string{"no", "n", "nah", "negative", "nope"}

func IsYesOrNo(txt string) (bool, error) {
	txt = strings.ToLower(strings.TrimSpace(txt))
	for _, v := range yes {
		if txt == v {
			return true, nil
		}
	}
	for _, v := range nos {
		if txt == v {
			return false, nil
		}
	}
	return false, errors.New("Not a yes or a no.")
}
