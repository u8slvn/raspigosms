package gsm

import (
	"regexp"
)

// CheckPhoneFormat test if the given string match with the E.164 phone format
// Todo : maybe use a more strict lib
func CheckPhoneFormat(phn string) bool {
	match, _ := regexp.MatchString("^\\+?[1-9]\\d{1,14}$", phn)
	return match
}
