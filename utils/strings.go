package utils

import (
	"unicode"
)

// Capitalize makes the first letter of the given text upper case
func Capitalize(text string) string {
	if len(text) == 0 {
		return ""
	}
	tmp := []rune(text)
	tmp[0] = unicode.ToUpper(tmp[0])
	return string(tmp)
}
