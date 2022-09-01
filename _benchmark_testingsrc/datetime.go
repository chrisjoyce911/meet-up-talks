package validator

import (
	"regexp"
	"time"
)

// DateRegex validates date format `yyyy-mm-dd`.
func DateRegex(v string) bool {
	return regexp.MustCompile(`([12]\d{3}-(0[1-9]|1[0-2])-(0[1-9]|[12]\d|3[01]))`).MatchString(v)
}

// TimeParse validates date format `yyyy-mm-dd`.
func TimeParse(v string) bool {
	if _, e := time.Parse("2006-01-02", v); e != nil {
		return false
	}
	return true
}
