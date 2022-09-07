package phone_normalizer

import (
	"regexp"
)

func PhoneNormalize(phone_number string) string {
	reg := regexp.MustCompile("\\D")
	return reg.ReplaceAllString(phone_number, "")
}
