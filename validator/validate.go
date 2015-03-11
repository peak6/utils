package validator

import (
	"github.com/plimble/utils/strings2"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"
)

var (
	emailPatern       = regexp.MustCompile(".+@.+\\..+")
	dateiso8601Patern = regexp.MustCompile("^(\\d{4})-(\\d{2})-(\\d{2})T(\\d{2}):(\\d{2}):(\\d{2})(Z|(\\+|-)\\d{2}(:?\\d{2})?)$")
)

func requiredString(val string) bool {
	if len(strings.TrimSpace(val)) == 0 {
		return true
	}

	return false
}

func requiredInt(val int) bool {
	if val == 0 {
		return true
	}

	return false
}

func requiredFloat64(val float64) bool {
	if val == 0 {
		return true
	}

	return false
}

func requiredBool(val bool) bool {
	if !val {
		return true
	}

	return false
}

func notNil(val interface{}) bool {
	if val == nil {
		return true
	}

	return false
}

func requiredTime(val time.Time) bool {
	if val.IsZero() {
		return true
	}

	return false
}

func minInt(val int, n int) bool {
	if val < n {
		return false
	}

	return true
}

func maxInt(val int, n int) bool {
	if val > n {
		return false
	}

	return true
}

func minFloat64(val float64, n float64) bool {
	if val < n {
		return false
	}

	return true
}

func maxFloat64(val float64, n float64) bool {
	if val > n {
		return false
	}

	return true
}

func minChar(val string, n int) bool {
	if utf8.RuneCount(string2.StringToBytes(val)) < n {
		return true
	}

	return false
}

func maxChar(val string, n int) bool {
	if utf8.RuneCount(string2.StringToBytes(val)) > n {
		return true
	}

	return false
}

func email(val string) bool {
	if val == "" {
		return false
	}
	if !emailPatern.MatchString(val) {
		return true
	}

	return false
}

func gender(val string) bool {
	if val != `male` && val != `female` {
		return true
	}

	return false
}

func confirmVals(val, confirm string) bool {
	if val != confirm {
		return true
	}

	return false
}

func iso8601DataTime(val string) bool {
	if val == "" {
		return false
	}
	if !dateiso8601Patern.MatchString(val) {
		return true
	}

	return false
}

func inString(val string, in []string) bool {
	for _, k := range in {
		if k == val {
			return false
		}
	}

	return true
}
