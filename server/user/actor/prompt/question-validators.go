package prompt

/*
	Validators for questions. Used by question builder.
*/
import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pwsdc/web-mud/util/language"
)

func intValidator(txt string) (bool, string) {
	txt = strings.TrimSpace(txt)
	_, err := strconv.Atoi(txt)
	if err != nil {
		return false, "You must answer with a whole number."
	} else {
		return true, ""
	}
}

func positiveIntValidator(txt string) (bool, string) {
	txt = strings.TrimSpace(txt)
	n, err := strconv.Atoi(txt)
	if err != nil {
		return false, "You must answer with a positive whole number."
	} else {
		if n < 1 {
			return false, "You must answer with a positive whole number."
		}
		return true, ""
	}
}

func intBetweenValidator(start int, end int) func(string) (bool, string) {
	too_low_msg := fmt.Sprintf("You must answer with a number greater than or equal to %d.", start)
	too_high_msg := fmt.Sprintf("You must answer with a number less than or equal to %d.", end)
	return func(txt string) (bool, string) {
		txt = strings.TrimSpace(txt)
		n, err := strconv.Atoi(txt)
		if err != nil {
			return false, "You must answer with a whole number."
		} else {
			if n < start {
				return false, too_low_msg
			}
			if n > end {
				return false, too_high_msg
			}
			return true, ""
		}

	}
}

func floatValidator(txt string) (bool, string) {
	txt = strings.TrimSpace(txt)
	_, err := strconv.ParseFloat(txt, 32)
	if err != nil {
		return false, "You must answer with a number."
	} else {
		return true, ""
	}
}

func floatBetweenValidator(start float32, end float32) func(string) (bool, string) {
	too_low_msg := fmt.Sprintf("You must answer with a number greater than or equal to %.2f.", start)
	too_high_msg := fmt.Sprintf("You must answer with a number less than or equal to %.2f.", end)
	return func(txt string) (bool, string) {
		txt = strings.TrimSpace(txt)
		n, err := strconv.ParseFloat(txt, 32)
		if err != nil {
			return false, "You must answer with a number."
		} else {
			n32 := float32(n)
			if n32 < start {
				return false, too_low_msg
			}
			if n32 > end {
				return false, too_high_msg
			}
			return true, ""
		}

	}
}

func yesNoValidator(txt string) (bool, string) {
	_, err := language.IsYesOrNo(txt)
	if err != nil {
		return false, "You must answer yes or no."
	}
	return true, ""
}

func multiChoiceValidator(choices []string) func(string) (bool, string) {
	not_in_answer := fmt.Sprintf("You must pick one of these answers: %s", strings.Join(choices, ", "))
	return func(txt string) (bool, string) {
		txt = strings.ToLower(strings.TrimSpace(txt))
		for _, v := range choices {
			if v == txt {
				return true, ""
			}
		}
		return false, not_in_answer
	}
}

func strlenValidator(length int) func(string) (bool, string) {
	bad := fmt.Sprintf("You must enter at least %d characters.", length)
	return func(txt string) (bool, string) {
		if len(txt) < length {
			return false, bad
		} else {
			return true, ""
		}
	}
}
