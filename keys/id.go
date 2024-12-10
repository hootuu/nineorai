package keys

import (
	"unicode"
)

type NineoraID string

func ValidateNineoraID(idStr string) bool {
	for _, ch := range idStr {
		if !unicode.IsDigit(ch) {
			return false
		}
	}
	return true
}
