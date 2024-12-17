package keys

import (
	"fmt"
	"strconv"
	"unicode"
)

type NineoraID string

func (n NineoraID) IsValid() bool {
	return ValidateNineoraID(string(n))
}

func ValidateNineoraID(idStr string) bool {
	for _, ch := range idStr {
		if !unicode.IsDigit(ch) {
			return false
		}
	}
	return true
}

func NineoraIdFromUint64(idUint64 uint64) NineoraID {
	str := strconv.FormatUint(idUint64, 10)
	length := len(str)
	targetLength := 4 + ((length-1)/2)*2
	format := fmt.Sprintf("%%0%dd", targetLength)
	return NineoraID(fmt.Sprintf(format, idUint64))
}
